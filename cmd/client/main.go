package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/AndrivA89/word-of-wisdom-pow/internal/pow"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatalf("Error reading from server: %v", err)
	}

	challengeMessage := string(buffer[:n])
	if !strings.HasPrefix(challengeMessage, "CHALLENGE") {
		log.Fatalf("Invalid challenge message from server")
	}

	parts := strings.Split(challengeMessage, " ")
	if len(parts) != 5 {
		log.Fatalf("Invalid challenge format")
	}

	// Extract the challenge string (everything except the first and last parts)
	challenge := strings.Join(parts[1:len(parts)-1], " ")

	// Extract the difficulty (the last part)
	difficultyStr := strings.TrimSpace(parts[len(parts)-1])
	difficulty, err := strconv.Atoi(difficultyStr)
	if err != nil {
		log.Fatalf("Error parsing difficulty: %v", err)
	}

	nonce := pow.SolveChallenge(challenge, difficulty)

	_, err = conn.Write([]byte(fmt.Sprintf("NONCE %s\n", nonce)))
	if err != nil {
		log.Fatalf("Error sending nonce to server: %v", err)
	}

	n, err = conn.Read(buffer)
	if err != nil {
		log.Fatalf("Error reading from server: %v", err)
	}

	fmt.Printf("Received citation: %s\n", string(buffer[:n]))
}
