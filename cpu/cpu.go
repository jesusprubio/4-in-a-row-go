package cpu

import (
	"4-in-a-row-go/game"
	"math/rand"
)

// ExecCPUTurn CPU's next move.
func ExecCPUTurn(b *game.Board) {
	x, y, sx := 0, 0, 0
	// Choose which side to start scanning on, left or right.
	if rand.Intn(2) == 1 {
		sx = 6
	}

	// find the place that can win.
	for x = 0; x < 7; x++ {
		if b.Height[x] > 9 {
			continue
		}
		y = 9 - b.Height[x]
		if checkCPUCell(x, y, 4, game.CPU, b) {
			b.Put(x, game.CPU)
			return
		}
	}

	// Prevent the player from making four in row.
	for x = 0; x < 7; x++ {
		if b.Height[x] > 9 {
			continue
		}
		y = 9 - b.Height[x]
		if checkCPUCell(x, y, 4, game.Player, b) {
			b.Put(x, game.CPU)
			return
		}
	}

	// Prevent the player from making three in row.
	for x = 0; x < 7; x++ {
		x2 := Abs(sx - x)
		if b.Height[x2] > 9 {
			continue
		}
		y = 9 - b.Height[x2]
		if checkCPUCell(x2, y, 3, game.Player, b) {
			// Avoiding the losing place of the next move.
			if IsCPULostNextTurn(x2, y, b) {
				continue
			}
			b.Put(x2, game.CPU)
			return
		}
	}

	// make three in row.
	for x = 0; x < 7; x++ {
		x2 := Abs(sx - x)
		if b.Height[x2] > 9 {
			continue
		}
		y = 9 - b.Height[x2]
		if checkCPUCell(x2, y, 3, game.CPU, b) {
			// Avoiding the losing place of the next move.
			if IsCPULostNextTurn(x2, y, b) {
				continue
			}
			b.Put(x2, game.CPU)
			return
		}
	}

	// Prevent the player from making two in row.
	for x = 0; x < 7; x++ {
		x2 := Abs(sx - x)
		if b.Height[x2] > 9 {
			continue
		}
		y = 9 - b.Height[x2]
		if checkCPUCell(x2, y, 2, game.Player, b) {
			// Avoiding the losing place of the next move.
			if IsCPULostNextTurn(x2, y, b) {
				continue
			}
			b.Put(x2, game.CPU)
			return
		}
	}

	// make two in row.
	/*
		for x = 0; x < 7; x++ {
			x2 := Abs(sx - x)
			if b.Height[x2] > 9 {
				continue
			}
			y = 9 - b.Height[x2]
			if checkCPUCell(x2, y, 2, game.CPU, b) {
				// Avoiding the losing place of the next move.
				if IsCPULostNextTurn(x2, y, b) {
					continue
				}
				b.Put(x2, game.CPU)
				return
			}
		}*/

	// random next move.
	for x = 0; x < 7; x++ {
		rx := rand.Intn(7)
		if b.Height[rx] > 9 {
			continue
		}
		y = 9 - b.Height[rx]
		// Avoiding the losing place of the next move.
		if IsCPULostNextTurn(rx, y, b) {
			continue
		}
		b.Put(rx, game.CPU)
		return
	}

	// if there is no place to expectrd, next move is an empty spot.
	for x = 0; x < 7; x++ {
		x2 := Abs(sx - x)
		if b.Height[x2] > 9 {
			continue
		}
		b.Put(x2, game.CPU)
		break
	}
}

// checkCPUCell　count the row.
func checkCPUCell(x, y, c int, z game.Char, b *game.Board) bool {
	cbord := b.Board
	cbord[y][x] = z
	return b.CheckCellCount(x, y, c, z, cbord)
}

// IsCPULostNextTurn Check if the CPU loses on the next move.
func IsCPULostNextTurn(x, y int, b *game.Board) bool {
	if b.Height[x] > 8 {
		return false
	}
	cbord := b.Board
	cbord[y][x] = game.CPU
	cbord[y-1][x] = game.Player

	return b.CheckCellCount(x, y-1, 4, game.Player, cbord)
}

// Abs
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
