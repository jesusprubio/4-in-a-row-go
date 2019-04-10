package cpu

import (
	"math/rand"

	"github.com/y-hatano-github/4inarow/game"
)

// ExecCPUTurn CPUのターン実行
func ExecCPUTurn(b *game.Board) {
	x, y, sx := 0, 0, 0
	// 左右どちらから走査するか
	if rand.Intn(2) == 1 {
		sx = 6
	}

	// 勝てる場所を探す
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

	// 負ける場所を探す
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

	// Playerが三つそろう箇所を抑える
	for x = 0; x < 7; x++ {
		x2 := Abs(sx - x)
		if b.Height[x2] > 9 {
			continue
		}
		y = 9 - b.Height[x2]
		if checkCPUCell(x2, y, 3, game.Player, b) {
			// 次の手で負ける場合は抑えない
			if IsCPULostNextTurn(x2, y, b) {
				continue
			}
			b.Put(x2, game.CPU)
			return
		}
	}

	// CPUが三つそろう箇所を抑える
	for x = 0; x < 7; x++ {
		x2 := Abs(sx - x)
		if b.Height[x2] > 9 {
			continue
		}
		y = 9 - b.Height[x2]
		if checkCPUCell(x2, y, 3, game.CPU, b) {
			// 次の手で負ける場合は抑えない
			if IsCPULostNextTurn(x2, y, b) {
				continue
			}
			b.Put(x2, game.CPU)
			return
		}
	}

	// Playerが二つそろう箇所を抑える
	for x = 0; x < 7; x++ {
		x2 := Abs(sx - x)
		if b.Height[x2] > 9 {
			continue
		}
		y = 9 - b.Height[x2]
		if checkCPUCell(x2, y, 2, game.Player, b) {
			// 次の手で負ける場合は抑えない
			if IsCPULostNextTurn(x2, y, b) {
				continue
			}
			b.Put(x2, game.CPU)
			return
		}
	}

	// CPUが二つそろう箇所を抑える
	/*
		for x = 0; x < 7; x++ {
			x2 := Abs(sx - x)
			if b.Height[x2] > 9 {
				continue
			}
			y = 9 - b.Height[x2]
			if checkCPUCell(x2, y, 2, game.CPU, b) {
				// 次の手で負ける場合は抑えない
				if IsCPULostNextTurn(x2, y, b) {
					continue
				}
				b.Put(x2, game.CPU)
				return
			}
		}*/

	// ランダムに手を置く
	for x = 0; x < 7; x++ {
		rx := rand.Intn(7)
		if b.Height[rx] > 9 {
			continue
		}
		y = 9 - b.Height[rx]
		// 負ける場所にはおかない
		if IsCPULostNextTurn(rx, y, b) {
			continue
		}
		b.Put(rx, game.CPU)
		return
	}

	// 期待した場所がなければ空いている場所に置く
	for x = 0; x < 7; x++ {
		x2 := Abs(sx - x)
		if b.Height[x2] > 9 {
			continue
		}
		b.Put(x2, game.CPU)
		break
	}
}

// checkCPUCell　CPUの手(期待値)のチェック
func checkCPUCell(x, y, c int, z game.Char, b *game.Board) bool {
	cbord := b.Board
	cbord[y][x] = z
	return b.CheckCellCount(x, y, c, z, cbord)
}

// IsCPULostNextTurn 次の手でCPUが負けるか
func IsCPULostNextTurn(x, y int, b *game.Board) bool {
	if b.Height[x] > 8 {
		return false
	}
	cbord := b.Board
	cbord[y][x] = game.CPU
	cbord[y-1][x] = game.Player

	return b.CheckCellCount(x, y-1, 4, game.Player, cbord)
}

// Abs 絶対値 mathだとfloatなので
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
