package main

import "fmt"

type Main struct {
	board [8][8]byte
	lines int
	cols int
}

func New() *Main{
	return &Main{lines: 8, cols: 8}
}

func (m *Main) Print() {
	for i := 0; i < m.lines; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Printf("%c", m.board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	main := New()
	main.Print()
}
