package main

import (
	"math/rand"
	"time"

	"4-in-a-row-go/cpu"
	"4-in-a-row-go/game"
	"4-in-a-row-go/player"
)

func main() {
	// 乱数
	rand.Seed(time.Now().UnixNano())

	b := &game.Board{}
	b.DrawTitle()
	for {
		// 初期化
		b.Init()
		b.DrawBoard()
		for {

			// プレイヤーのターン
			player.ExecPlayerTurn(b)
			// 盤描画
			b.DrawBoard()
			// チェック
			b.Judge()
			if b.GameStatus != game.Playing {
				break
			}
			// COMのターン
			cpu.ExecCPUTurn(b)
			// 盤描画
			b.DrawBoard()
			// チェック
			b.Judge()
			if b.GameStatus != game.Playing {
				break
			}
		}
		// 終了処理
		b.EndGame()
	}
}
