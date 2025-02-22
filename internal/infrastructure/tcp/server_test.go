package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/AndrivA89/word-of-wisdom-pow/internal/application/mocks"
	"github.com/AndrivA89/word-of-wisdom-pow/internal/domain"
	"github.com/AndrivA89/word-of-wisdom-pow/internal/pow"
)

// findValidNonce computes a valid nonce for the given challenge and difficulty using the actual pow.VerifySolution.
func findValidNonce(challenge string, difficulty int) string {
	nonce := 0
	for {
		candidate := fmt.Sprintf("%d", nonce)
		if pow.VerifySolution(challenge, candidate, difficulty) {
			return candidate
		}
		nonce++
	}
}

// TestHandleConnection_CorrectSolution verifies that a correct solution yields a citation.
func TestHandleConnection_CorrectSolution(t *testing.T) {
	expectedCitation := "Test citation"
	expectedAuthor := "Test Author"
	citation := domain.NewCitation(expectedCitation, expectedAuthor)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mocks.NewMockcitationRepository(ctrl)
	mockService.EXPECT().GetRandomCitation().Return(citation, nil).Times(1)

	server := NewServer("0.0.0.0:9000", mockService, 4, 5*time.Second)

	clientConn, serverConn := net.Pipe()
	defer func(clientConn net.Conn) {
		assert.NoError(t, clientConn.Close())
	}(clientConn)
	defer func(serverConn net.Conn) {
		assert.NoError(t, serverConn.Close())
	}(serverConn)

	go server.handleConnection(serverConn)

	reader := bufio.NewReader(clientConn)
	challengeLine, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Error reading challenge: %v", err)
	}
	challengeLine = strings.TrimSpace(challengeLine)
	parts := strings.Split(challengeLine, " ")
	if len(parts) < 3 {
		t.Fatalf("Invalid challenge format: %s", challengeLine)
	}
	// Expected format: "CHALLENGE <challenge> <difficulty>"
	// The challenge text may contain spaces; join all parts except the first and last.
	challengeText := strings.Join(parts[1:len(parts)-1], " ")
	diff, err := strconv.Atoi(strings.TrimSpace(parts[len(parts)-1]))
	if err != nil {
		t.Fatalf("Error parsing difficulty: %v", err)
	}

	// Compute a valid nonce using the actual pow.VerifySolution.
	validNonce := findValidNonce(challengeText, diff)

	_, err = clientConn.Write([]byte("NONCE " + validNonce + "\n"))
	if err != nil {
		t.Fatalf("Error sending nonce: %v", err)
	}

	response, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Error reading citation: %v", err)
	}

	response = strings.TrimSpace(response)
	if !strings.HasPrefix(response, "CITATION:") {
		t.Fatalf("Expected citation response, got: %s", response)
	}
	if !strings.Contains(response, expectedCitation) || !strings.Contains(response, expectedAuthor) {
		t.Fatalf("Response does not contain expected citation and author. Got: %s", response)
	}
}
