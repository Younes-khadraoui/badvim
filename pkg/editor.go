package editor

import (
	"bufio"
	"fmt"
	"os"
)

type Editor struct {
	filePath      string
	content       []string
	newLineBuffer string
	mode          string
}

func NewEditor(filePath string) (*Editor, error) {
	var content []string

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &Editor{
		filePath: filePath,
		content:  content,
		mode:     "normal",
	}, nil
}

func (e *Editor) Run() error {
	for {
		if e.mode == "normal" {
			fmt.Println("Normal Mode:")
			for _, line := range e.content {
				fmt.Println(line)
			}
			fmt.Print("Command (i to enter input mode, : to execute command): ")

			key, err := readKey()
			if err != nil {
				return err
			}

			switch key {
			case 'i':
				e.mode = "input"
			case ':':
				fmt.Print("'q' to quit: ")
				command, err := readLine()
				if err != nil {
					return err
				}
				if command == "q" {
					if len(e.newLineBuffer) > 0 {
						if err := os.WriteFile(e.filePath, []byte(e.newLineBuffer), 0644); err != nil {
							return err
						}
					}
					fmt.Println("Exiting BadVim. Goodbye!")
					return nil
				} else {
					fmt.Printf("Command '%s' is not recognized.\n", command)
				}
			default:
				fmt.Println("Invalid command. Press 'i' to enter input mode or ':' to execute command.")
			}
		} else if e.mode == "input" {
			fmt.Print("Input Mode (Press 'ESC' to return to normal mode): ")

			for {
				key, err := readKey()
				if err != nil {
					return err
				}

				if key == 27 { // ESC key
					e.mode = "normal"
					fmt.Println("Returning to Normal Mode.")
					break
				} else {
					e.newLineBuffer += string(key)
				}
			}
		}
	}
}

func readKey() (byte, error) {
	var b [1]byte
	_, err := os.Stdin.Read(b[:])
	if err != nil {
		return 0, err
	}
	return b[0], nil
}

func readLine() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", scanner.Err()
}
