package main

import (
	"log"

	"github.com/tg-checker/internal/gateway/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("run gateway app: %v", err)
	}
}
