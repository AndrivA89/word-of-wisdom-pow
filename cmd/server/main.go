package main

import (
	"log"

	"github.com/AndrivA89/word-of-wisdom-pow/internal/infrastructure/repository"
	"github.com/AndrivA89/word-of-wisdom-pow/internal/infrastructure/tcp"
)

func main() {
	repo, err := repository.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	server := tcp.NewServer(":9000", repo, 4)
	server.Start()
}
