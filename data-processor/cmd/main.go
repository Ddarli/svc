package main

import (
	"data-processor/config"
	storage2 "data-processor/internal/storage"
)

func main() {
	storage := storage2.New(config.Config{})

	_ = storage
}
