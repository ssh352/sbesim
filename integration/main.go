package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"sbe/configure"
	"sbe/entity"
	fix "sbe/sbe/iLinkBinary"
	"sbe/service"
	"strconv"
	"time"
)

func Send(con net.Conn, msg entity.SBEMessage) error {
	min := fix.NewSbeGoMarshaller()
	var buf = new(bytes.Buffer)
	if err := msg.Encode(min, buf, true); err != nil {
		fmt.Println("Encoding Error", err)
		return err
	}
	hdr := fix.MessageHeader{
		BlockLength: msg.SbeBlockLength(),
		TemplateId:  msg.SbeTemplateId(),
		SchemaId:    msg.SbeSchemaId(),
		Version:     msg.SbeSchemaVersion(),
	}
	var hdrBuf = new(bytes.Buffer)
	if err := hdr.Encode(min, hdrBuf); err != nil {
		fmt.Println("Encoding Error", err)
		return err
	}
	hdrBytes := hdrBuf.Bytes()
	bodyBytes := buf.Bytes()
	frame := make([]byte, 4)
	binary.LittleEndian.PutUint16(frame[0:], uint16(len(hdrBytes))+uint16(len(bodyBytes)))
	binary.LittleEndian.PutUint16(frame[2:], 0xCAFE)
	log.Printf("write length %d for frame body 2 bytes", len(hdrBytes)+len(bodyBytes))
	log.Printf("write endian %d 2 bytes", 0xCAFE)
	_, err := con.Write(frame)
	if err != nil {
		return err
	}
	var n int
	n, err = con.Write(hdrBytes)
	if err != nil {
		return err
	}
	log.Printf("write %d bytes for header", n)
	n, err = con.Write(bodyBytes)
	if err == nil {
		log.Printf("write %d bytes for body", n)
	}
	return err
}

func randStr(len int) string {
	res := ""
	for i := 0; i < len; i++ {
		res = res + strconv.Itoa(rand.Int()%10)
	}
	return res
}

func RunIntTest(ctx context.Context, conn net.Conn) error {
	log.Println("sending negotiate msg")

	negotiate := fix.Negotiate500{}
	negotiate.UUID = rand.Uint64()
	negotiate.RequestTimestamp = uint64(time.Now().UTC().UnixNano())
	copy(negotiate.AccessKeyID[:], randStr(20))
	copy(negotiate.Firm[:], randStr(5))
	copy(negotiate.HMACSignature[:], randStr(32))
	copy(negotiate.Session[:], randStr(3))
	if err := Send(conn, &negotiate); err != nil {
		return err
	}
	log.Println("negotiate msg sent")

	m := fix.NewSbeGoMarshaller()
	bt := make([]byte, 4)
	for {
		_, err := io.ReadFull(conn, bt)
		if err != nil {
			log.Printf("failed to read: %d dummy bytes", len(bt))
			return err
		}
		var hdr fix.SbeGoMessageHeader
		if err := hdr.Decode(m, conn); err != nil {
			log.Printf("failed to read header due to %v", err)
			return err
		}
		templateID := hdr.TemplateId
		msg, err := service.CreateMsg(int(templateID))
		if err != nil {
			log.Printf("failed to create msg due to %v", err)
			return err
		}
		if err := msg.Decode(m, conn, hdr.Version, hdr.BlockLength, false); err != nil {
			log.Println("Failed to decode msg", err)
			return err
		}
		if templateID == 501 { // negotiate done
			log.Println("sending establish msg")
			establish := fix.Establish503{}
			establish.Session = negotiate.Session
			establish.HMACSignature = negotiate.HMACSignature
			establish.Firm = negotiate.Firm
			establish.AccessKeyID = negotiate.AccessKeyID
			establish.RequestTimestamp = uint64(time.Now().UTC().UnixNano())
			establish.NextSeqNo = 1
			establish.UUID = negotiate.UUID
			establish.KeepAliveInterval = 5
			copy(establish.TradingSystemName[:], randStr(30))
			copy(establish.TradingSystemVendor[:], randStr(10))
			copy(establish.TradingSystemVersion[:], randStr(10))
			if err = Send(conn, &establish); err != nil {
				return err
			}
			log.Println("establish msg send")
		}
		if templateID == 504 { // establish ack
			log.Println("sending new order")
			order := fix.NewOrderSingle514{}
			order.SeqNum = 1
			order.OrderQty = rand.Uint32()
			order.ManualOrderIndicator = fix.ManualOrdIndReq.Automated
			order.TimeInForce = fix.TimeInForce.Day
			order.Side = fix.SideReq.Buy
			order.OrdType = fix.OrderTypeReq.Limit
			order.ExpireDate = 321
			order.SecurityID = 123
			order.SendingTimeEpoch = rand.Uint64()
			order.Price.Mantissa = rand.Int63()
			order.Price.Exponent = 1
			order.StopPx.Mantissa = rand.Int63()
			order.StopPx.Exponent = 2
			order.PartyDetailsListReqID = rand.Uint64()
			copy(order.ClOrdID[:], randStr(20))
			copy(order.Location[:], "US/IL")
			copy(order.SenderID[:], randStr(20))
			order.LiquidityFlag = fix.BooleanNULL.False
			order.MinQty = rand.Uint32()
			if err = Send(conn, &order); err != nil {
				return err
			}
			log.Println("new order sent")
		}
		if templateID == 525 {
			log.Println("fill received")
			conn.Close()
			return nil
		}
	}
	return nil
}

func main() {
	conf := configure.New()
	addr := fmt.Sprintf("127.0.0.1:%s", conf.Port)
	log.Println("connecting to " + addr)
	var con net.Conn
	var err error
	for i := 0; i < 20; i++ {
		con, err = net.Dial("tcp", addr)
		if err != nil {
			time.Sleep(time.Second * 1)
			log.Printf("Retry %d time", i)
			continue
		}
		break
	}
	if con == nil {
		fmt.Println("Failed to create client connection")
		os.Exit(-1)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = RunIntTest(ctx, con)
	code := 0
	if err != nil {
		code = 1
		fmt.Printf("failure: %s", err.Error())
	} else {
		log.Println("Integration passed")
	}
	os.Exit(code)
}
