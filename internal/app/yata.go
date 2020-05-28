//Package app private application and library code
package app

import (
	"bufio"
	"fmt"
	"os"
)

// Yata Structure for yata application
type Yata struct {
	Todolists []*TodoList
}

// Read all TodoLists from a file. Only supports one list
// Pointers vs Values param/return: https://stackoverflow.com/questions/23542989/pointers-vs-values-in-parameters-and-return-values/23551970#23551970
func Read(filename string) TodoList {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening: ", err)
		panic("Error while opening the file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(bufio.NewReader(file))

	var todos []Todo
	for scanner.Scan() {
		todoName := scanner.Text()
		item := Todo{Title: todoName, Done: false}
		todos = append(todos, item)
	}
	todolist := TodoList{Name: "Main", Todo: todos}
	return todolist
}

// Write all TodoList to file. For the moment supports only one list
func Write(filename string, todolist []TodoList) {
	file, err := os.Create(filename)
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	tl := todolist[0]
	for _, item := range tl.Todo {
		fmt.Fprintf(w, "%s\n", item.Title)
	}
}

// Display a list of todos
func Display() {
	todolist := Read("/Users/gosselinchristian/tmp/yata.txt")
	fmt.Printf("%s \n", todolist.Name)
	for j, item := range todolist.Todo {
		fmt.Printf("\t%d: %s\n", j, item.Title)
	}
}
