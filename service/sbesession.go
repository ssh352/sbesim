package service

import (
	"net"
)

type sbeSession struct {
	conn   net.Conn
}

func (s *sbeSession) Serve() error {
	return nil
}