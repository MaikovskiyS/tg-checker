package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/tg-checker/internal/checker/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("run checker app: %v", err)
	}
}
