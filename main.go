package main

import (
	"fmt"
	"log"

	"github.com/alwindoss/vector/internal/server"
)

func main() {
	fmt.Println("Vector: Railways Management System")
	log.Fatal(server.Run())
}
