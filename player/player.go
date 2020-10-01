package player

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"4-in-a-row-go/game"
)

// ExecPlayerTurn Player's next move.
func ExecPlayerTurn(b *game.Board) {
	for {
		fmt.Print("type number 1-9. (type 'q' to quit):")
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		t := stdin.Text()
		if t == "q" {
			fmt.Println("bye..")
			os.Exit(0)
		}
		if i, err := strconv.Atoi(t); err == nil {
			if i > 0 && i < 8 {
				if b.Put(i-1, 1) {
					break
				}
			}
		}
	}
}
