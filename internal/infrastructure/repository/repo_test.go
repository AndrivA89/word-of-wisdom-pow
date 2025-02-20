package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomCitation(t *testing.T) {
	repo, err := NewRepository()
	assert.NoError(t, err)

	citation, err := repo.GetRandomCitation()
	assert.NoError(t, err)
	assert.NotEmpty(t, citation)
}
