package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

const (
	ColorReset   = "\033[0m"
	ColorReverse = "\033[7m"
)

type Main struct {
	board    [][]rune
	count    [][]int
	overall  int
	useColor bool
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
		'B': {
			single:     false,
			reverse:    false,
			directions: []Direction{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}},
		},
		'b': {
			single:     false,
			reverse:    true,
			directions: []Direction{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}},
		},
		'N': {
			single:     true,
			reverse:    false,
			directions: []Direction{{-2, -1}, {-1, -2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1}},
		},
		'n': {
			single:     true,
			reverse:    true,
			directions: []Direction{{-2, -1}, {-1, -2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1}},
		},
		'Q': {
			single:     false,
			reverse:    false,
			directions: []Direction{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}, {-1, 0}, {1, 0}, {0, -1}, {0, 1}},
		},
		'q': {
			single:     false,
			reverse:    true,
			directions: []Direction{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}, {-1, 0}, {1, 0}, {0, -1}, {0, 1}},
		},
		'K': {
			single:     true,
			reverse:    false,
			directions: []Direction{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}, {-1, 0}, {1, 0}, {0, -1}, {0, 1}},
		},
		'k': {
			single:     true,
			reverse:    true,
			directions: []Direction{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}, {-1, 0}, {1, 0}, {0, -1}, {0, 1}},
		},
	}
)

func New(useColor bool) *Main {
	m := &Main{useColor: useColor, board: [][]rune{
		[]rune("rnbqkbnr"),
		[]rune("pppppppp"),
		[]rune("        "),
		[]rune("        "),
		[]rune("        "),
		[]rune("        "),
		[]rune("PPPPPPPP"),
		[]rune("RNBQKBNR"),
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
			if m.useColor && (i+j)%2 == 0 {
				fmt.Print(ColorReverse)
			}
			fmt.Printf("%c ", v)
			if m.useColor {
				fmt.Print(ColorReset)
			}
		}
		fmt.Println()
		for j, _ := range row {
			if m.useColor && (i+j)%2 == 0 {
				fmt.Print(ColorReverse)
			}
			v := m.count[i][j]
			if v != 0 {
				fmt.Printf("%2d", v)
			} else {
				fmt.Printf("  ")
			}
			if m.useColor {
				fmt.Print(ColorReset)
			}
		}
		fmt.Println()
	}
	fmt.Printf("overall %d\n", m.overall)
}

func (m *Main) PrintEachPiece(waitForUserInput bool) {
	for i, row := range m.board {
		for j, _ := range row {
			m.board[i][j] = ' '
		}
	}
	m.Update()
	m.Print()
	mi := len(m.board) / 2
	mj := len(m.board[0]) / 2
	for v, _ := range Pieces {
		m.board[mi][mj] = v
		m.Update()
		m.Print()
		m.board[mi][mj] = ' '

		if waitForUserInput {
			buf := bufio.NewReader(os.Stdin)
			fmt.Print("> ")
			if _, err := buf.ReadBytes('\n'); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func main() {
	useColor := flag.Bool("use_color", true, "use color")
	doPrintEachPiece := flag.Bool("print_each_piece", false, "print each piece")
	flag.Parse()

	main := New(*useColor)
	main.Print()
	if *doPrintEachPiece {
		main.PrintEachPiece(*useColor)
	}
}
