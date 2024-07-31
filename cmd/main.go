package main

import (
	"fmt"
	"log"
	"os"

	"github.com/younes-khadraoui/badvim/"
)

func main() {
	// Get file path from command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Error: No file path provided. Usage: badvim <file>")
		return
	}
	filePath := os.Args[1]

	// Initialize editor with file path
	e, err := editor.NewEditor(filePath)
	if err != nil {
		log.Fatalf("Error initializing editor: %v", err)
	}

	fmt.Println("Welcome to BadVim - Press 'i' to enter input mode, 'ESC' to exit input mode.")

	oldState, err := utils.SetRawMode()
	if err != nil {
		fmt.Println("Error setting terminal mode:", err)
		return
	}
	defer utils.RestoreMode(oldState)

	if err := e.Run(); err != nil {
		fmt.Println("Error running editor:", err)
	}
}
