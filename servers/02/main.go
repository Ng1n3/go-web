package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Server struct {
	listener net.Listener
	clients  sync.Map
	logger   *log.Logger
}

func NewServer(addr string) (*Server, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Server{
		listener: listener,
		logger:   log.New(os.Stdout, "TCP Server: ", log.Ldate|log.Ltime|log.Lshortfile),
		clients:  sync.Map{},
	}, nil
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Set connection deadline
	conn.SetDeadline(time.Now().Add(10 * time.Minute))

	// Store client connection
	s.clients.Store(conn.RemoteAddr(), conn)
	defer s.clients.Delete(conn.RemoteAddr())

	s.logger.Printf("New connection from %s", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		s.logger.Printf("Received from %s: %s", conn.RemoteAddr(), message)

		// Echo back with timestamp
		response := fmt.Sprintf("[%s] Server received: %s", 
			time.Now().Format(time.RFC3339), message)
		
		conn.Write([]byte(response + "\n"))

		// Optional: Special command handling
		if message == "quit" {
			conn.Write([]byte("Goodbye!\n"))
			break
		}
	}

	if err := scanner.Err(); err != nil {
		s.logger.Printf("Error reading from %s: %v", conn.RemoteAddr(), err)
	}
}

func (s *Server) Start() error {
	s.logger.Println("Server starting on", s.listener.Addr())

	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				s.logger.Printf("Accept error: %v", err)
				continue
			}
			go s.handleConnection(conn)
		}
	}()

	// Graceful shutdown handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	return s.Stop()
}

func (s *Server) Stop() error {
	s.logger.Println("Shutting down server...")

	// Close all active client connections
	s.clients.Range(func(key, value interface{}) bool {
		if conn, ok := value.(net.Conn); ok {
			conn.Close()
		}
		return true
	})

	return s.listener.Close()
}

func main() {
	server, err := NewServer(":3030")
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}