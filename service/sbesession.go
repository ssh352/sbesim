package service

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"sbe/entity"
	"sbe/handler"
	fix "sbe/sbe/iLinkBinary"
)

type sbeSession struct {
	Conn  net.Conn
	h     handler.OrderEntry
	sUUID uint64
	seqNo uint32
}

func NewSBEServerSession(c net.Conn, h handler.OrderEntry) entity.SBESession {
	return &sbeSession{
		Conn:  c,
		h:     h,
		sUUID: 0,
		seqNo: 0,
	}
}

func (s *sbeSession) GetSeqNo() uint32 {
	return s.seqNo
}
func (s *sbeSession) SetSeqNo(v uint32) {
	s.seqNo = v
}
func (s *sbeSession) GetUUID() uint64 {
	return s.sUUID
}

func (s *sbeSession) Send(msg entity.SBEMessage) error {
	min := fix.NewSbeGoMarshaller()
	var buf = new(bytes.Buffer)
	if err := msg.Encode(min, buf, true); err != nil {
		fmt.Println("Encoding Error", err)
		return err
	}
	hdr := fix.MessageHeader{
		BlockLength: msg.SbeBlockLength(),
		TemplateId: msg.SbeTemplateId(),
		SchemaId: msg.SbeSchemaId(),
		Version : msg.SbeSchemaVersion(),
	}
	var hdrBuf =  new(bytes.Buffer)
	if err := hdr.Encode(min, hdrBuf); err != nil {
		fmt.Println("Encoding Error", err)
		return err
	}
	hdrBytes := hdrBuf.Bytes()
	bodyBytes := buf.Bytes()
	frame := make([]byte, 4)
	binary.LittleEndian.PutUint16(frame[0:], uint16(len(hdrBytes)) + uint16(len(bodyBytes)))
	binary.LittleEndian.PutUint16(frame[2:],  0xCAFE)
	_, err := s.Conn.Write(frame)
	if err != nil {
		_, err = s.Conn.Write(hdrBytes)
	}
	if err != nil {
		_, err = s.Conn.Write(bodyBytes)
	}
	return err
}

func (s *sbeSession) Close() error {
	return s.Conn.Close()
}

func (s *sbeSession) Serve() error {
	m := fix.NewSbeGoMarshaller()
	bt := make([]byte, 4)
	for {
		_, err := io.ReadFull(s.Conn, bt)
		if err != nil {
			log.Printf("failed to read: %d dummy bytes", len(bt))
			return err
		}
		var hdr fix.SbeGoMessageHeader
		if err := hdr.Decode(m, s.Conn); err != nil {
			log.Printf("failed to read header due to %v", err)
			return err
		}
		templateID := hdr.TemplateId
		msg, err := createMsg(int(templateID))
		if err != nil {
			log.Printf("failed to create msg due to %v", err)
			return err
		}
		if err := msg.Decode(m, s.Conn, hdr.Version, hdr.BlockLength, true); err != nil {
			log.Println("Failed to decode msg", err)
			return err
		}

		// only process what we interested
		err = nil
		switch templateID {
		case 500:
			err = s.onNegotiate(msg.(*fix.Negotiate500))
		case 503:
			err = s.onEstablish(msg.(*fix.Establish503))
		case 506:
			err = s.onSequence(msg.(*fix.Sequence506))
		case 507:
			err = s.onTerminate(msg.(*fix.Terminate507))
		case 508:
			err = s.onRetransmissionRequest(msg.(*fix.RetransmitRequest508))
		case 514:
			s.h.OnOrderNew(context.Background(), msg.(*fix.NewOrderSingle514), s)
		case 516:
			s.h.OnOrderCancel(context.Background(), msg.(*fix.OrderCancelRequest516), s)
		case 515:
			s.h.OnOrderModify(context.Background(), msg.(*fix.OrderCancelReplaceRequest515), s)
		}
		if err != nil {
			log.Println("Failed to process msg", err)
			return err
		}
	}
	return nil
}

func (s *sbeSession) onNegotiate(msg *fix.Negotiate500) error {
	// always ack
	resp := fix.NegotiationResponse501{}
	resp.UUID = msg.UUID
	s.sUUID = msg.UUID
	resp.FaultToleranceIndicator = fix.FTI.Primary
	resp.SecretKeySecureIDExpiration = 100
	resp.RequestTimestamp = msg.RequestTimestamp
	resp.PreviousSeqNo = 0
	resp.PreviousUUID = 0
	return s.Send(&resp)
}

func (s *sbeSession) onEstablish(msg *fix.Establish503) error {
	resp := fix.EstablishmentAck504{}
	resp.UUID = msg.UUID
	resp.RequestTimestamp = msg.RequestTimestamp
	resp.NextSeqNo = s.seqNo + 1
	s.seqNo = s.seqNo + 1
	resp.PreviousSeqNo = 0
	resp.PreviousUUID = 0
	resp.KeepAliveInterval = msg.KeepAliveInterval
	resp.FaultToleranceIndicator = fix.FTI.Primary
	return s.Send(&resp)
}

func (s *sbeSession) onSequence(msg *fix.Sequence506) error {
	resp := fix.Sequence506{}
	resp.FaultToleranceIndicator = fix.FTI.Primary
	resp.NextSeqNo = s.seqNo + 1
	resp.UUID = msg.UUID
	resp.KeepAliveIntervalLapsed = msg.KeepAliveIntervalLapsed
	return s.Send(&resp)
}

func (s *sbeSession) onTerminate(msg *fix.Terminate507) error {
	resp := fix.Terminate507{}
	resp.UUID = msg.UUID
	resp.RequestTimestamp = msg.RequestTimestamp
	resp.ErrorCodes = 0
	return s.Send(&resp)
}

func (s *sbeSession) onRetransmissionRequest(msg *fix.RetransmitRequest508) error {
	resp := fix.RetransmitReject510{}
	resp.UUID = msg.UUID
	resp.ErrorCodes = 0
	return s.Send(&resp)
}
