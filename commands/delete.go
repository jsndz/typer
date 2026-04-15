package commands

import (
	"log"
	"os"
)

func Delete(filename string) {
	err := os.Remove(filename)
	if err != nil {
		log.Fatal("error in removing")
	}
}
