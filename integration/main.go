package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	faker "github.com/bxcodec/faker"
	"io"
	"log"
	"net"
	"os"
	"sbe/configure"
	"sbe/entity"
	fix "sbe/sbe/iLinkBinary"
	"sbe/service"
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
	_, err := con.Write(frame)
	if err != nil {
		_, err = con.Write(hdrBytes)
	}
	if err != nil {
		_, err = con.Write(bodyBytes)
	}
	return err
}

func RunIntTest(ctx context.Context, conn net.Conn) error {
	log.Println("sending negotiate msg")

	negotiate := fix.Negotiate500{}
	err := faker.FakeData(&negotiate)
	if err != nil {
		return err
	}
	if err = Send(conn, &negotiate); err != nil {
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
		if err := msg.Decode(m, conn, hdr.Version, hdr.BlockLength, true); err != nil {
			log.Println("Failed to decode msg", err)
			return err
		}
		if templateID == 501 { // negotiate done
			log.Println("sending establish msg")
			establish := fix.Establish503{}
			err = faker.FakeData(&establish)
			if err != nil {
				return err
			}
			if err = Send(conn, &establish); err != nil {
				return err
			}
			log.Println("establish msg send")
		}
		if templateID == 504 { // establish ack
			log.Println("sending new order")
			neworder := fix.NewOrderSingle514{}
			err = faker.FakeData(&neworder)
			if err != nil {
				return err
			}
			if err = Send(conn, &neworder); err != nil {
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
	}
	os.Exit(code)
}
