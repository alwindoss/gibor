package main

import (
	"fmt"
	"log"

	"github.com/alwindoss/gibor/internal/server"
)

func main() {
	fmt.Println("Welcome to Gibor")
	cfg := server.Config{}
	log.Fatal(server.Run(&cfg))
}
