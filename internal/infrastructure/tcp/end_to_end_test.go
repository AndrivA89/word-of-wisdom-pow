package tcp

import (
	"bufio"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/AndrivA89/word-of-wisdom-pow/internal/application/mocks"
	"github.com/AndrivA89/word-of-wisdom-pow/internal/domain"
)

// TestEndToEnd performs an end-to-end test of the TCP server and client flow.
func TestEndToEnd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Expected citation to be returned by the repository.
	expectedCitation := "Integration Test Citation"
	expectedAuthor := "Integration Author"
	citation := domain.NewCitation(expectedCitation, expectedAuthor)

	// Create a mock citation repository that returns our citation.
	mockRepo := mocks.NewMockcitationRepository(ctrl)
	mockRepo.EXPECT().GetRandomCitation().Return(citation, nil).Times(1)

	// Start server on a random available port.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	assert.NoError(t, err)

	defer func(ln net.Listener) {
		assert.NoError(t, ln.Close())
	}(ln)
	actualAddr := ln.Addr().String()

	server := NewServer(actualAddr, mockRepo, 4, 5*time.Second)

	// Launch server in a separate goroutine.
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				return
			}
			go server.handleConnection(conn)
		}
	}()

	// Give the server a moment to start.
	time.Sleep(500 * time.Millisecond)

	conn, err := net.Dial("tcp", actualAddr)
	assert.NoError(t, err)

	defer func(conn net.Conn) {
		assert.NoError(t, conn.Close())
	}(conn)

	reader := bufio.NewReader(conn)
	// Read the challenge message from the server.
	challengeLine, err := reader.ReadString('\n')
	assert.NoError(t, err)
	challengeLine = strings.TrimSpace(challengeLine)
	parts := strings.Split(challengeLine, " ")

	// Expected format: "CHALLENGE <challenge text> <difficulty>"
	challengeText := strings.Join(parts[1:len(parts)-1], " ")
	diff, err := strconv.Atoi(strings.TrimSpace(parts[len(parts)-1]))
	assert.NoError(t, err)

	// Compute a valid nonce using the actual pow.VerifySolution.
	validNonce := findValidNonce(challengeText, diff)

	// Client sends the solution.
	_, err = conn.Write([]byte("NONCE " + validNonce + "\n"))
	assert.NoError(t, err)

	// Read the response (citation) from the server.
	response, err := reader.ReadString('\n')
	assert.NoError(t, err)

	response = strings.TrimSpace(response)
	assert.True(t, strings.HasPrefix(response, "CITATION:"))
	assert.True(t, strings.Contains(response, expectedCitation) || !strings.Contains(response, expectedAuthor))
}
