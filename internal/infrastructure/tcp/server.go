package tcp

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/AndrivA89/word-of-wisdom-pow/internal/infrastructure/repository"
)

type Server struct {
	address    string
	repository *repository.Repo
}

func NewServer(address string, repository *repository.Repo) *Server {
	return &Server{address: address, repository: repository}
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
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	_, err := conn.Write([]byte("Welcome! Request a quote by typing 'GET QUOTE'.\n"))
	if err != nil {
		log.Printf("Error sending message: %v", err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading client data: %v", err)
		return
	}

	message := strings.TrimSpace(string(buffer[:n]))

	if strings.ToUpper(message) == "GET QUOTE" {
		citation, err := s.repository.GetRandomCitation()
		if err != nil {
			conn.Write([]byte(fmt.Sprintf("Error getting citation: %v\n", err)))
			return
		}

		// Send the citation to the client
		conn.Write([]byte(fmt.Sprintf("CITATION: %s\n", citation)))
	} else {
		conn.Write([]byte("To get a quote, type 'GET QUOTE'.\n"))
	}
}
