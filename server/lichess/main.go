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

func MapMoves(board *yacboard.YacBoard) (stats, bool) {
	r := explorer.Request(board.GetFEN(), speeds.SlowerThanBullet, ratings.All)
	//

	totalGames := 0
	for _, v := range r.Moves {
		totalGames += v.White + v.Draws + v.Black
	}
	//fmt.Print("games: ", totalGames)
	if totalGames < 100 {
		return stats{r.White, r.Draws, r.Black}, false
	}
	selectRandom := rand.Intn(totalGames)
	for _, v := range r.Moves {
		selectRandom -= v.White + v.Draws + v.Black
		if selectRandom <= 0 {
			//fmt.Println(", selected: ", v.San)
			if !board.DoUciMove(v.Uci) {
				panic("move not found: " + v.San + " (" + v.Uci + ")")
			}
			return stats{v.White, v.Draws, v.Black}, true
		}
	}
	panic("selected move not found")
}

func Main() {
	rand.Seed(time.Now().UnixNano())
	max := 0
	maxFast := 0
	for {
		board := yacboard.New()
		board.DoUciMove("e2e4", "f7f5")
		//fmt.Println(board)
		tim := time.Now()
		for i := 0; i < max; i++ {
			_, next := MapMoves(&board)
			if !next {
				break
			}
		}
		if time.Since(tim).Milliseconds() < 100 {
			maxFast++
			if maxFast > 100 {
				max++
				maxFast = 0
				fmt.Println("--- Next Depth ", max, " ---")
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}
}
