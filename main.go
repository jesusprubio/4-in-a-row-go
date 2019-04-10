package main

import (
	"math/rand"
	"time"

	"github.com/y-hatano-github/4inarow/cpu"
	"github.com/y-hatano-github/4inarow/game"
	"github.com/y-hatano-github/4inarow/player"
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
			player.ExecPlayerTunr(b)
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
