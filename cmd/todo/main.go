package main

import (
	"fmt"
	"os"
	"strings"

	todo "github.com/ezzy77/command-go"
)

const todoFileName = ".todo.json"

func main() {

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// decide what to do based on the number of arguments provided

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)

		}
	default:
		// concatenate all args with a space
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)

		// save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

}
