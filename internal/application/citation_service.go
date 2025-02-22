// Package application service for repository
package application

import (
	"github.com/AndrivA89/word-of-wisdom-pow/internal/domain"
)

// CitationService - service struct.
type CitationService struct {
	repo citationRepository
}

// NewCitationService - create service with repo.
func NewCitationService(repo citationRepository) *CitationService {
	return &CitationService{repo: repo}
}

// GetRandomCitation - get random citation from repo.
func (s *CitationService) GetRandomCitation() (*domain.Citation, error) {
	return s.repo.GetRandomCitation()
}
