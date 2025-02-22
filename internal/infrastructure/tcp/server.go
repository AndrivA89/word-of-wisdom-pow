package tcp

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/AndrivA89/word-of-wisdom-pow/internal/pow"
)

// Server - struct for config server.
type Server struct {
	address    string
	service    citationService
	difficulty int
	timeout    time.Duration
}

// NewServer - create new server.
func NewServer(address string, service citationService, difficulty int, timeout time.Duration) *Server {
	return &Server{address: address, service: service, difficulty: difficulty, timeout: timeout}
}

// Start - starting server.
func (s *Server) Start() {
	ln, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer func(ln net.Listener) {
		err = ln.Close()
		if err != nil {
			log.Printf("Error closing server: %v", err)
		}
	}(ln)

	log.Printf("Server is listening on %s", s.address)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		log.Printf("Connection established with %s", conn.RemoteAddr())

		err = conn.SetDeadline(time.Now().Add(s.timeout)) // Set timeout for PoW solving
		if err != nil {
			log.Printf("Error set deadline: %v", err)
			return
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}(conn)

	challenge := pow.GenerateChallenge()
	log.Printf("Generated challenge: %s", challenge)

	if err := writeMessage(conn, fmt.Sprintf("CHALLENGE %s %d\n", challenge, s.difficulty)); err != nil {
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading client data: %v", err)
		return
	}

	message := strings.TrimSpace(string(buffer[:n]))
	log.Printf("Received solution from client: %s", message)

	parts := strings.Split(message, " ")
	if len(parts) != 2 || parts[0] != "NONCE" {
		if err = writeMessage(conn, "Invalid solution format.\n"); err != nil {
			return
		}
		return
	}

	nonce := parts[1]
	if !pow.VerifySolution(challenge, nonce, s.difficulty) {
		if err = writeMessage(conn, "Incorrect solution.\n"); err != nil {
			return
		}
		return
	}

	citation, err := s.service.GetRandomCitation()
	if err != nil {
		if err = writeMessage(conn, fmt.Sprintf("Error getting citation: %v\n", err)); err != nil {
			return
		}
		return
	}

	if err = writeMessage(
		conn,
		fmt.Sprintf("CITATION: %s AUTHOR: %s\n",
			citation.GetText(),
			citation.GetAuthor(),
		),
	); err != nil {
		return
	}

	log.Printf("Sent citation to client: %s, Author: %s",
		citation.GetText(), citation.GetAuthor())
}

func writeMessage(conn net.Conn, message string) error {
	if _, err := conn.Write([]byte(message)); err != nil {
		log.Printf("error write data to the connection: %s", err)
		return err
	}

	return nil
}
