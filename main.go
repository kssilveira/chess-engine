package main

import "fmt"

type Main struct {
	board []string
}

func New() *Main{
	return &Main{board: []string{
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
	}}
}

func (m *Main) Print() {
	for _, row := range m.board {
		fmt.Printf("%s\n", row)
	}
	fmt.Println()
}

func main() {
	main := New()
	main.Print()
}
