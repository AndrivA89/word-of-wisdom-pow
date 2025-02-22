// Package tcp - interface for server
package tcp

import "github.com/AndrivA89/word-of-wisdom-pow/internal/domain"

type citationService interface {
	GetRandomCitation() (*domain.Citation, error)
}
