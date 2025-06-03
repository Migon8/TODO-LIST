package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func main()  {

	const (

		Create = "Create"
		Read = "Read"
		Update = "Update"
		Delete = "Delete"

	)


	fmt.Println("Welcome to the TO DO List CLI app")
	fmt.Println()
	fmt.Println("Enter your command (Create, Read, Update, Delete):")


	reader := bufio.NewReader(os.Stdin)

	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command)


	if command == Create {
		fmt.Println("Enter your task:")
		task, _ := reader.ReadString('\n')
		task = strings.TrimSpace(task)

		file, err := os.OpenFile("TODO-LIST", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
		defer file.Close()

		file.WriteString(task + "\n")
		fmt.Println("Задача добавлена!")
	}
 
	if command == Read {
	
		readFile, err := os.Open("TODO-LIST")
		if err !=nil {
			fmt.Println("Ошибка при чтении файла:", err)
			return
		}
		defer readFile.Close()

		scanner :=bufio.NewScanner(readFile)
		fmt.Println("Ваш список задач:")
		for scanner.Scan() {
			fmt.Println("-" + scanner.Text())
		}

		if err :=scanner.Err(); err !=nil {
			fmt.Println("Ошибка чтения")
		}
	}



	

}