// Package main - entry point for server
package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/AndrivA89/word-of-wisdom-pow/internal/application"
	"github.com/AndrivA89/word-of-wisdom-pow/internal/infrastructure/repository"
	"github.com/AndrivA89/word-of-wisdom-pow/internal/infrastructure/tcp"
)

func main() {
	repo, err := repository.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	difficulty := 1
	envDifficulty := os.Getenv("DIFFICULTY")
	if envDifficulty != "" {
		difficulty, err = strconv.Atoi(envDifficulty)
		if err != nil {
			log.Fatal(err)
		}
	}

	service := application.NewCitationService(repo)

	server := tcp.NewServer(":9000", service, difficulty, time.Second)
	server.Start()
}
