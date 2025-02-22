package application

import "github.com/AndrivA89/word-of-wisdom-pow/internal/domain"

type citationRepository interface {
	GetRandomCitation() (*domain.Citation, error)
}
