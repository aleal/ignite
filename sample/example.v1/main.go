package main

import (
	"context"
	"log"

	"github.com/aleal/ignite/example.v1"
	"github.com/americanas-go/config"
)

func main() {
	config.Load()
	instance, e := example.New(context.Background())
	if e != nil {
		log.Fatalf("Unexpected error %v", e)
	}

	log.Printf("New Instance of %T -> %v", instance, instance)
}
