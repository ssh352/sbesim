package service

import (
	"context"
	"io"
	"log"
	"net"
	"sbe/handler"
	fix "sbe/sbe/iLinkBinary"
)

type sbeSession struct {
	Conn net.Conn
	h    handler.OrderEntry
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
			s.h.OnOrderNew(context.Background(), msg.(*fix.NewOrderSingle514), &s.Conn)
		case 516:
			s.h.OnOrderCancel(context.Background(), msg.(*fix.OrderCancelRequest516), &s.Conn)
		case 515:
			s.h.OnOrderModify(context.Background(), msg.(*fix.OrderCancelReplaceRequest515), &s.Conn)
		}
		if err != nil {
			log.Println("Failed to process msg", err)
			return err
		}
	}
	return nil
}

func (s *sbeSession) onNegotiate(msg *fix.Negotiate500) error {
	return nil
}

func (s *sbeSession) onEstablish(msg *fix.Establish503) error {
	return nil
}

func (s *sbeSession) onSequence(msg *fix.Sequence506) error {
	return nil
}

func (s *sbeSession) onTerminate(msg *fix.Terminate507) error {
	return nil
}

func (s *sbeSession) onRetransmissionRequest(msg *fix.RetransmitRequest508) error {
	return nil
}
