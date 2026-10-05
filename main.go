package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"example.com/notetaker/note"
)

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if input == "" {
		return "", errors.New("input cannot be empty")
	}

	return input, nil
}

func getNoteDetails() (string, string, error) {
	title, err := getUserInput("Enter the title of the note: ")
	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Enter the content of the note: ")
	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

func main() {
	//Take Title and Content from user for Note
	title, content, err := getNoteDetails()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//create a new Note
	note, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//Store it in a file
	err = note.SaveToFile()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
