package main

import "fmt"

const (
	ColorReset   = "\033[0m"
	ColorReverse = "\033[7m"
)

type Main struct {
	board   []string
	count   [][]int
	overall int
}

type Piece struct {
	single     bool
	value      int
	reverse    bool
	directions []Direction
}

type Direction struct {
	i int
	j int
}

var (
	Pieces = map[rune]Piece{
		'P': {
			single:     true,
			reverse:    false,
			directions: []Direction{{-1, -1}, {-1, 1}},
		},
		'p': {
			single:     true,
			reverse:    true,
			directions: []Direction{{-1, -1}, {-1, 1}},
		},
		'R': {
			single:     false,
			reverse:    false,
			directions: []Direction{{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
		},
		'r': {
			single:     false,
			reverse:    true,
			directions: []Direction{{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
		},
	}
)

func New() *Main {
	m := &Main{board: []string{
		"rnbqkbnr",
		"pppppppp",
		"        ",
		"        ",
		"        ",
		"        ",
		"PPPPPPPP",
		"RNBQKBNR",
	}}
	m.count = make([][]int, len(m.board))
	for i, row := range m.board {
		m.count[i] = make([]int, len(row))
	}
	m.Update()
	return m
}

func (m *Main) Update() {
	for i, row := range m.board {
		for j, _ := range row {
			m.count[i][j] = 0
		}
	}

	for i, row := range m.board {
		for j, v := range row {
			piece, ok := Pieces[v]
			if !ok {
				continue
			}
			reverse := 1
			if piece.reverse {
				reverse = -1
			}
			for _, direction := range piece.directions {
				for delta := 1; ; delta++ {
					ni := i + direction.i*delta*reverse
					nj := j + direction.j*delta*reverse
					if ni < 0 || ni >= len(m.board) || nj < 0 || nj >= len(row) {
						break
					}
					m.count[ni][nj] += reverse
					if piece.single || m.board[ni][nj] != ' ' {
						break
					}
				}
			}
		}
	}

	m.overall = 0
	for _, row := range m.count {
		for _, v := range row {
			if v > 0 {
				m.overall += 1
			} else if v < 0 {
				m.overall -= 1
			}
		}
	}
}

func (m *Main) Print() {
	for i, row := range m.board {
		for j, v := range row {
			if (i+j)%2 == 0 {
				fmt.Print(ColorReverse)
			}
			fmt.Printf("%c ", v)
			fmt.Print(ColorReset)
		}
		fmt.Println()
		for j, _ := range row {
			if (i+j)%2 == 0 {
				fmt.Print(ColorReverse)
			}
			fmt.Printf("%2d", m.count[i][j])
			fmt.Print(ColorReset)
		}
		fmt.Println()
	}
	fmt.Printf("overall %d\n", m.overall)
}

func main() {
	main := New()
	main.Print()
}
