package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	const (
		Create = "Create"
		Read   = "Read"
		Update = "Update"
		Delete = "Delete"
	)

	fmt.Println("Welcome to the TO DO List CLI app")
	fmt.Println()
	fmt.Println("Enter your command (Create, Read, Update, Delete):")

	reader := bufio.NewReader(os.Stdin)

	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command)

	switch command {

	case Create:
		fmt.Println("Enter your task:")
		task, _ := reader.ReadString('\n')
		task = strings.TrimSpace(task)

		file, err := os.OpenFile("TODO-LIST", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error while opening file for writing:", err)
			return
		}

		defer file.Close()

		_, err = file.WriteString(task + "\n")
		if err != nil {
			fmt.Println("Error while writing task to file:", err)
			return
		}
		fmt.Println("Task added!")

	case Read:
		file, err := os.Open("TODO-LIST")
		if err != nil {
			fmt.Println("Error while opening file for reading:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		fmt.Println("Your task list:")
		for scanner.Scan() {
			fmt.Println("- " + scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error while reading the file:", err)
		}

	case Update:
		fmt.Println("Update functionality is not implemented yet.")

	case Delete:
		fmt.Println("Update functionality is not implemented yet.")

	default:
		fmt.Println("Unknown command:", command)

	}

}
