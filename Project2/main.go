package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/interface/todo"
)

func main() {
	// title, content := getNoteDate()
	// userNote, err := note.New(title, content)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// userNote.Display()
	// err = userNote.Save()

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Saving the note succeeded!")

	content := getTodoDate()
	userTodo, err := todo.New(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userTodo.Display()
	err = userTodo.Save()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Saving the note succeeded!")

}

func getTodoDate() string {
	content := getUserInput("Note content:")
	return content
}

func getNoteDate() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content:")
	return title, content
}

// helper function to get user input
func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	// Scanf will take a single input
	// var value string
	// fmt.Scan(&value)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")

	return text
}
