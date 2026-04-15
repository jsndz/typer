package commands

import (
	"bufio"
	"log"
	"os"
)

func Create(filename string) {
	// file offset moves to end of file cause of os.0_APPEND
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("open error:", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		_, err = file.Write(append(scanner.Bytes(), '\n'))
		if err != nil {
			log.Fatal("write error:", err)
		}
	}
}
