package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

func Create(filename string) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	undoStack := ([][]byte{})
	// redoStack := ([][]byte{})
	var current []byte

	// file offset moves to end of file cause of os.0_APPEND
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("open error:", err)
	}
	defer file.Close()
	reader := bufio.NewReader(os.Stdin)

	// file.Read(current)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			log.Fatal(err)
		}
		switch b {
		case 3:
			fmt.Println("\nExit")
			return
		case 19:
			err := file.Truncate(0)
			if err != nil {
				log.Fatal(err)
			}

			_, err = file.Seek(0, 0)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Saving: " + string(current))
			file.Write((current))
		case 10:
			undoStack = append(undoStack, append([]byte{}, current...))
			current = append(current, '\n')
			fmt.Printf("\n")
		default:
			undoStack = append(undoStack, append([]byte{}, current...))
			current = append(current, b)
			fmt.Printf("%c", b)
		}

		if err != nil {
			log.Fatal("write error:", err)
		}
	}
}
