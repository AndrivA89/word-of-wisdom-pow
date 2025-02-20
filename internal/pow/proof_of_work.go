package pow

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func GenerateChallenge() string {
	return "Solve this challenge"
}

func VerifySolution(challenge string, nonce string, difficulty int) bool {
	hashInput := challenge + nonce
	hash := sha256.Sum256([]byte(hashInput))
	hashHex := fmt.Sprintf("%x", hash)

	return strings.HasPrefix(hashHex, strings.Repeat("0", difficulty))
}

func SolveChallenge(challenge string, difficulty int) string {
	nonce := 0
	for {
		hashInput := challenge + fmt.Sprintf("%d", nonce)
		hash := sha256.Sum256([]byte(hashInput))
		hashHex := fmt.Sprintf("%x", hash)

		// Check if the hash starts with the required number of leading zeros
		if strings.HasPrefix(hashHex, strings.Repeat("0", difficulty)) {
			break
		}
		nonce++
	}
	return fmt.Sprintf("%d", nonce)
}
