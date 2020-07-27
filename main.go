package main

import (
	"log"

	"github.com/lughong/blog-service/internal/router"
)

func main() {
	r := router.NewRouter()

	addr := ":8080"
	if err := r.Run(addr); err != nil {
		log.Fatalf("router.Run error. %s", err)
	}
}
