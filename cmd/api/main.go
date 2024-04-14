package main

import (
	"github.com/daniskazan/avito-tech-banner-service/internal/handlers"
)

func main() {
	srv := handlers.NewServer()
	srv.Start()
}
