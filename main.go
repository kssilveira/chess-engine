package main

import "fmt"

const (
	ColorReset   = "\033[0m"
	ColorReverse = "\033[7m"
)

type Main struct {
	board []string
}

func New() *Main {
	return &Main{board: []string{
		"rnbqkbnr",
		"pppppppp",
		"        ",
		"        ",
		"        ",
		"        ",
		"PPPPPPPP",
		"RNBQKBNR",
	}}
}

func (m *Main) Print() {
	for i, row := range m.board {
		for j, col := range row {
			if (i+j)%2 == 0 {
				fmt.Print(ColorReverse)
			}
			fmt.Printf("%c", col)
			fmt.Print(ColorReset)
		}
		fmt.Println()
	}
}

func main() {
	main := New()
	main.Print()
}
