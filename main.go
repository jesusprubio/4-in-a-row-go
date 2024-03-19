package main

import (
	"math/rand"
	"time"

	"github.com/jesusprubio/4-in-a-row-go/ai"
	"github.com/jesusprubio/4-in-a-row-go/game"
	"github.com/jesusprubio/4-in-a-row-go/player"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	b := &game.Board{}
	b.DrawTitle()
	for {
		// init
		b.Init()
		b.DrawBoard()
		for {

			// player's turn
			player.ExecPlayerTurn(b)
			// draw bord
			b.DrawBoard()
			// check
			b.Judge()
			if b.GameStatus != game.Playing {
				break
			}
			// AI's turn
			ai.ExecCPUTurn(b)
			// draw bord
			b.DrawBoard()
			// check
			b.Judge()
			if b.GameStatus != game.Playing {
				break
			}
		}
		// draw bord
		b.DrawBoard()

		// game end
		b.EndGame()
	}
}
