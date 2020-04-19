package service

import (
	"context"
	"log"
	"net"
	"sbe/entity"
	"sbe/handler"
	"sync"
)

// SBEServer ...
type SBEServer interface {
	Listen(ctx context.Context, address string) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

// New ...
func New(oeh handler.OrderEntry) SBEServer {
	return &sbeServer{
		oe: oeh,
	}
}

type sbeServer struct {
	oe handler.OrderEntry
	listener net.Listener
	sessions []entity.SBESession
	mutex   *sync.Mutex
}

func (s *sbeServer) Listen(ctx context.Context, address string) error {
	l, err := net.Listen("tcp", address)
	if err == nil {
		s.listener = l
	}
	log.Printf("Listening on %v", address)
	return err
}

func (s *sbeServer) Start(ctx context.Context)  error{
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Print(err)
		} else {
			// handle connection
			client := s.accept(conn)
			go s.serve(client)
		}
	}
	return nil
}

func (s *sbeServer) Stop(ctx context.Context) error {
	for _, s := range s.sessions {
		s.Close()
	}
	s.sessions = nil
	s.listener.Close()
	return nil
}

func (s *sbeServer) accept(conn net.Conn) entity.SBESession {
	log.Printf("Accepting connection from %v, total clients: %v", conn.RemoteAddr().String(), len(s.sessions)+1)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	client := NewSBEServerSession(conn, s.oe)
	s.sessions = append(s.sessions, client)
	return client
}
func (s *sbeServer) remove(client entity.SBESession) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// remove the connections from clients array
	for i, check := range s.sessions {
		if check == client {
			s.sessions = append(s.sessions[:i], s.sessions[i+1:]...)
		}
	}
	log.Printf("Closing connection ")
	client.Close()
}

func (s *sbeServer) serve(client entity.SBESession) {
	defer s.remove(client)
	client.Serve()
}