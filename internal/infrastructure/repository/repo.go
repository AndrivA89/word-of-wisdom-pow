// Package repository - load citations
package repository

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/AndrivA89/word-of-wisdom-pow/internal/domain"
)

const fileName = "citations.txt"

// Repo - struct for repository.
type Repo struct {
	citations []*domain.Citation
}

// NewRepository - create new repository, load citation from file.
func NewRepository() (*Repo, error) {
	filePath, err := getFilePath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open the file: %w", err)
	}

	citations := make([]*domain.Citation, 0, 101)
	scanner := bufio.NewScanner(file)
	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading the file: %w", err)
	}

	re := regexp.MustCompile(`^"(.*?)"\s*—\s*(.*)$`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) == 3 {
			// matches[1] — citation, matches[2] — author
			citations = append(citations, domain.NewCitation(matches[1], matches[2]))
		}
	}

	if len(citations) == 0 {
		return nil, fmt.Errorf("no citations found in the file")
	}

	if err = file.Close(); err != nil {
		return nil, err
	}

	return &Repo{citations: citations}, nil
}

// GetRandomCitation - get random citation from repo.
func (r *Repo) GetRandomCitation() (*domain.Citation, error) {
	index, err := randomIndex(len(r.citations))
	if err != nil {
		return nil, err
	}

	return r.citations[index], nil
}

func randomIndex(limit int) (int, error) {
	randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(limit)))
	return int(randomNum.Int64()), err
}

func getFilePath() (string, error) {
	filePath := os.Getenv("CITATIONS_FILE_PATH")

	if filePath == "" {
		wd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("failed to get current directory: %w", err)
		}

		rootRepoPath, _, found := strings.Cut(wd, "word-of-wisdom-pow")
		if !found {
			return "", err
		}

		filePath = filepath.Join(
			rootRepoPath,
			"word-of-wisdom-pow",
			"internal",
			"infrastructure",
			"repository",
			fileName,
		)
	}

	return filePath, nil
}
