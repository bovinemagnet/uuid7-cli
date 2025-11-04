package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

func main() {
	// Generate UUIDv7 based on current time
	id, err := uuid.NewV7()
	if err != nil {
		fmt.Printf("Error generating UUIDv7: %v\n", err)
		os.Exit(1)
	}

	// Print the UUID to screen
	fmt.Println(id.String())
}
