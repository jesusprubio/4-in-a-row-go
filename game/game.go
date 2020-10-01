package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shiena/ansicolor"
)

// Char 表示キャラクタ制御
type Char int

// 表示キャラクタ
const (
	Brank Char = iota
	Player
	CPU
)

// Status 状態制御
type Status int

// status
const (
	Playing Status = iota
	PlayerWin
	CPUWin
	Draw
)

// Board
type Board struct {
	// 盤(縦×横 : 10 × 7)
	Board [10][7]Char
	// 盤の埋まっている高さ
	Height [7]int
	// ゲーム状態
	GameStatus Status
}

// Init
func (b *Board) Init() {
	for i, rows := range b.Board {
		for j := range rows {
			b.Board[i][j] = 0
			b.Height[j] = 0
		}
	}
	b.GameStatus = Playing
}

// DrawTitle
func (b *Board) DrawTitle() {
	w := ansicolor.NewAnsiColorWriter(os.Stdout)
	fmt.Println("4 in a row.")
	fmt.Fprint(w, fmt.Sprintf("\x1b[32m%s\x1b[0m", "〇"))
	fmt.Println(" is your cell.")
	fmt.Fprint(w, fmt.Sprintf("\x1b[31m%s\x1b[0m", "〇"))
	fmt.Println(" is cpu's cell.")
	for {
		fmt.Print("type 's' to start(type 'q' to quit):")
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		t := stdin.Text()

		if t == "q" {
			fmt.Println("bye..")
			os.Exit(0)
		}
		if t == "s" {
			break
		}
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Game start!!")
	fmt.Println("")
}

// EndGame
func (b *Board) EndGame() {
	switch b.GameStatus {
	case PlayerWin:
		fmt.Println("Player won!!")
	case CPUWin:
		fmt.Println("Player lost..")
	case Draw:
		fmt.Println("Draw game.")
	}

	for {
		fmt.Print("type 'r' to restart(type 'q' to quit):")
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		t := stdin.Text()

		if t == "q" {
			fmt.Println("bye..")
			os.Exit(0)
		}
		if t == "r" {
			fmt.Println()
			fmt.Println()
			fmt.Println()
			break
		}
	}
}

// DrawBoard
func (b *Board) DrawBoard() {
	for _, rows := range b.Board {
		var a string
		for _, value := range rows {
			switch value {
			case Brank:
				a += "・"
			case Player:
				a += fmt.Sprintf("\x1b[32m%s\x1b[0m", "〇")
			case CPU:
				a += fmt.Sprintf("\x1b[31m%s\x1b[0m", "〇")
			}
		}
		w := ansicolor.NewAnsiColorWriter(os.Stdout)
		fmt.Fprintln(w, a)
	}
	fmt.Println("１２３４５６７")
	fmt.Println()
}

// Put
func (b *Board) Put(x int, z Char) bool {
	if b.Height[x] == 10 {
		fmt.Printf("[%v] not vacant.", x+1)
		fmt.Println()
		return false
	}
	b.Board[9-b.Height[x]][x] = z
	b.Height[x]++
	return true
}

//Judge
func (b *Board) Judge() {
	for y, rows := range b.Board {
		for x, value := range rows {
			if value == 0 {
				continue
			}
			if b.CheckCellCount(x, y, 4, value, b.Board) {
				switch value {
				case Player:
					b.GameStatus = PlayerWin
					return
				case CPU:
					b.GameStatus = CPUWin
					return
				}
			}
		}
	}

	if b.IsDraw() {
		b.GameStatus = Draw
		return
	}
	b.GameStatus = Playing
}

// IsDraw
func (b *Board) IsDraw() bool {
	for _, value := range b.Height {
		if value < 10 {
			return false
		}
	}

	return true
}

// CheckCellCount
func (b *Board) CheckCellCount(x, y, c int, z Char, board [10][7]Char) bool {
	cs := ""
	for i := 0; i < c; i++ {
		cs += [...]string{"0", "1", "2"}[z]
	}

	// column
	cells := ""
	for yy := y - (c - 1); yy < y+c; yy++ {
		if yy > -1 && yy < 10 {
			cells += [...]string{"0", "1", "2"}[board[yy][x]]
		}
	}
	if strings.Index(cells, cs) != -1 {
		return true
	}

	// row
	cells = ""
	for xx := x - (c - 1); xx < x+c; xx++ {
		if xx > -1 && xx < 7 {
			cells += [...]string{"0", "1", "2"}[board[y][xx]]
		}
	}
	if strings.Index(cells, cs) != -1 {
		return true
	}

	// Right shoulder down.
	cells = ""
	xx := x - (c - 1)
	for yy := y - (c - 1); yy < y+c; yy++ {
		if (yy > -1 && yy < 10) && (xx > -1 && xx < 7) {
			cells += [...]string{"0", "1", "2"}[board[yy][xx]]
		}
		xx++
	}
	if strings.Index(cells, cs) != -1 {
		return true
	}

	// Left shoulder down
	cells = ""
	xx = x + (c - 1)
	for yy := y - (c - 1); yy < y+c; yy++ {
		if (yy > -1 && yy < 10) && (xx > -1 && xx < 7) {
			cells += [...]string{"0", "1", "2"}[board[yy][xx]]
		}
		xx--
	}
	if strings.Index(cells, cs) != -1 {
		return true
	}

	return false
}
