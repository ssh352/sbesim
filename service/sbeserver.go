package service

import (
	"context"
	"log"
	"net"
	"sync"
)

// SBEServer ...
type SBEServer interface {
	Listen(ctx context.Context, address string) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}



type sbeServer struct {
	listener net.Listener
	sessions []*sbeSession
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

func (s *sbeServer) Start(ctx context.Context) {
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
}

func (s *sbeServer) accept(conn net.Conn) *sbeSession {
	log.Printf("Accepting connection from %v, total clients: %v", conn.RemoteAddr().String(), len(s.sessions)+1)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	client := &sbeSession{
		conn:   conn,
	}
	s.sessions = append(s.sessions, client)
	return client
}
func (s *sbeServer) remove(client *sbeSession) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// remove the connections from clients array
	for i, check := range s.sessions {
		if check == client {
			s.sessions = append(s.sessions[:i], s.sessions[i+1:]...)
		}
	}
	log.Printf("Closing connection from %v", client.conn.RemoteAddr().String())
	client.conn.Close()
}

func (s *sbeServer) serve(client *sbeSession) {
	defer s.remove(client)
	for {
		if err :=  client.Serve(); err != nil {
			log.Printf("Closing connection due to %s", err.Error())
			break
		}
	}
}