package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	list := []string{}
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a response: ")
	for scanner.Scan() {
		input = scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input")
		}
		if strings.Contains(input, "add") {
			var index int = strings.Index(input, "add")

			addTodo(&list, input[(index+3):])
		} else if strings.Contains(input, "done") {
			var index int = strings.Index(input, "done")
			doneItem, err := strconv.Atoi(input[index+5:])
			if err != nil {
				fmt.Println("error converting done item")
			}
			fmt.Println(input[index+4:])
			finishedToDo(&list, doneItem)
		} else if strings.Contains(input, "list") {
			fmt.Println(list)
		}
	}
	fmt.Println(list)

}

func addTodo(list *[]string, task string) {
	*list = append(*list, task)

}
func finishedToDo(list *[]string, item int) {
	*list = slices.Delete(*list, item+1, item+2)
}
