package book

import (
	"fmt"
	"strings"
)

func (book Book) PrintTodos() error {
	for _, day := range book.Days {
		for _, entry := range day.Entries {
			todos := readTodos(entry.Text)
			if len(todos) > 0 {
				fmt.Printf("%s %02d/%02d/%d\n", day.Date.Weekday().String(), day.Date.Day(), int(day.Date.Month()), day.Date.Year())
				for _, todo := range todos {
					fmt.Printf("%s\n", todo)
				}
			}
		}
	}

	return nil
}

func readTodos(text string) []string {
	todos := []string{}

	lines := strings.Split(text, "\n")

	start := -1
	for i, line := range lines {
		switch line {
		case "```TODO":
			start = i + 1
		case "```":
			if start != -1 {
				todos = append(todos, strings.Join(lines[start:i], "\n"))
				start = -1
			}
		}
	}

	return todos
}
