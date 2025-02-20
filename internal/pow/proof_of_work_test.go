package pow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoW(t *testing.T) {
	challenge := GenerateChallenge()
	nonce := SolveChallenge(challenge, 6)

	assert.True(t, VerifySolution(challenge, nonce, 6))
}
