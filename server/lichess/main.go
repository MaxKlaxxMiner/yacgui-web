package lichess

import (
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard"
	"math/rand"
	"server/lichess/explorer"
	"server/lichess/explorer/ratings"
	"server/lichess/explorer/speeds"
	"time"
)

type stats struct {
	white int
	draws int
	black int
}

func MapMoves(board *yacboard.YacBoard) bool {
	r := explorer.Request(board.GetFEN(), speeds.SlowerThanBullet, ratings.All)
	//

	totalGames := 0
	for _, v := range r.Moves {
		totalGames += v.White + v.Draws + v.Black
	}
	fmt.Print("games: ", totalGames)
	if totalGames < 100 {
		return false
	}
	selectRandom := rand.Intn(totalGames)
	for _, v := range r.Moves {
		selectRandom -= v.White + v.Draws + v.Black
		if selectRandom <= 0 {
			fmt.Println(", selected: ", v.San)
			moves := board.GetMoves()
			if v.San == "O-O" || v.San == "O-O-O" {
				for _, m := range moves {
					if m.Uci() == "e1g1" && board.WhiteMove && v.San == "O-O" {
						board.DoMove(m)
						return true
					}
					if m.Uci() == "e1c1" && board.WhiteMove && v.San == "O-O-O" {
						board.DoMove(m)
						return true
					}
					if m.Uci() == "e8g8" && !board.WhiteMove && v.San == "O-O" {
						board.DoMove(m)
						return true
					}
					if m.Uci() == "e8c8" && !board.WhiteMove && v.San == "O-O-O" {
						board.DoMove(m)
						return true
					}
				}
			}
			for _, m := range moves {
				if m.Uci() == v.Uci {
					board.DoMove(m)
					return true
				}
			}
			panic("move not found: " + v.San + " (" + v.Uci + ")")
		}
	}
	panic("selected move not found")
}

func Main() {
	rand.Seed(time.Now().UnixNano())
	board := yacboard.New()
	fmt.Println(board)
	for MapMoves(&board) {
	}
}
