package tcp

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/AndrivA89/word-of-wisdom-pow/internal/infrastructure/repository"
	"github.com/AndrivA89/word-of-wisdom-pow/internal/pow"
)

type Server struct {
	address    string
	repository *repository.Repo
	difficulty int
}

func NewServer(address string, repo *repository.Repo, difficulty int) *Server {
	return &Server{address: address, repository: repo, difficulty: difficulty}
}

func (s *Server) Start() {
	ln, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer ln.Close()

	log.Printf("Server is listening on %s", s.address)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		log.Printf("Connection established with %s", conn.RemoteAddr())
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	challenge := pow.GenerateChallenge()
	log.Printf("Generated challenge: %s", challenge)

	_, err := conn.Write([]byte(fmt.Sprintf("CHALLENGE %s %d\n", challenge, s.difficulty)))
	if err != nil {
		log.Printf("Error sending challenge: %v", err)
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
		conn.Write([]byte("Invalid solution format.\n"))
		return
	}

	nonce := parts[1]
	if !pow.VerifySolution(challenge, nonce, s.difficulty) {
		conn.Write([]byte("Incorrect solution.\n"))
		return
	}

	citation, err := s.repository.GetRandomCitation()
	if err != nil {
		conn.Write([]byte(fmt.Sprintf("Error getting citation: %v\n", err)))
		return
	}

	conn.Write([]byte(fmt.Sprintf("CITATION: %s\n", citation)))
	log.Printf("Sent citation to client: %s", citation)
}
