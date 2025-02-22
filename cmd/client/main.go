// Package main - entry point for client
package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/AndrivA89/word-of-wisdom-pow/internal/pow"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer func(conn net.Conn) {
		err = conn.Close()
		if err != nil {
			log.Printf("Error closing server: %v", err)
		}
	}(conn)

	err = conn.SetDeadline(time.Now().Add(10 * time.Second)) // Set timeout for PoW solving
	if err != nil {
		log.Printf("Error set deadline: %v", err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatalf("Error reading from server: %v", err)
	}

	nonce := solveChallenge(buffer, n)
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

func solveChallenge(buffer []byte, n int) string {
	challengeMessage := string(buffer[:n])
	if !strings.HasPrefix(challengeMessage, "CHALLENGE") {
		log.Fatalf("Invalid challenge message from server")
	}

	parts := strings.Split(challengeMessage, " ")
	challenge := strings.Join(parts[1:len(parts)-1], " ")

	// Extract the difficulty (the last part)
	difficultyStr := strings.TrimSpace(parts[len(parts)-1])
	difficulty, err := strconv.Atoi(difficultyStr)
	if err != nil {
		log.Fatalf("Error parsing difficulty: %v", err)
	}

	return pow.SolveChallenge(challenge, difficulty)
}
