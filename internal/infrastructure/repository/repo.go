package repository

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
)

const fileName = "citations.txt"

type Repo struct {
	citations []string
}

func NewRepository() (*Repo, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current directory: %w", err)
	}

	rootRepoPath, _, found := strings.Cut(wd, "word-of-wisdom-pow")
	if !found {
		return nil, err
	}

	filePath := filepath.Join(
		rootRepoPath,
		"word-of-wisdom-pow",
		"internal",
		"infrastructure",
		"repository",
		fileName,
	)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open the file: %w", err)
	}

	citations := make([]string, 0, 101)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		citations = append(citations, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading the file: %w", err)
	}

	if len(citations) == 0 {
		return nil, fmt.Errorf("no citations found in the file")
	}

	if err = file.Close(); err != nil {
		return nil, err
	}

	return &Repo{citations: citations}, nil
}

func (r *Repo) GetRandomCitation() (string, error) {
	index, err := randomIndex(len(r.citations))
	if err != nil {
		return "", err
	}

	return r.citations[index], nil
}

func randomIndex(limit int) (int, error) {
	randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(limit)))
	return int(randomNum.Int64()), err
}
