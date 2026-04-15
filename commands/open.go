package commands

import (
	"io"
	"log"
	"os"
)

func Open(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Error in opening file")
	}
	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		log.Fatalf("failed to stream file: %s", err)
	}
}
