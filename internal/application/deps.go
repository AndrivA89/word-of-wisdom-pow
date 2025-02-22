package application

import "github.com/AndrivA89/word-of-wisdom-pow/internal/domain"

//go:generate mockgen -source=deps.go -destination=mocks/mock.go -package=mocks

type citationRepository interface {
	GetRandomCitation() (*domain.Citation, error)
}
