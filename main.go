package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Print("Hello world")
	fmt.Printf("UUDI: %s", id)
}
