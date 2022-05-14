package main

import (
	"context"
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/YacBoard"
	"log"
	"mime"
	"net/http"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)

func wsReader(c *websocket.Conn, r *http.Request) {
	for {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		var v any
		err := wsjson.Read(ctx, c, &v)
		cancel()
		if err != nil {
			if websocket.CloseStatus(err) != websocket.StatusNormalClosure &&
				websocket.CloseStatus(err) != websocket.StatusGoingAway {
				log.Println(err)
			}
			return
		}

		fmt.Println("received:", v)
	}
}

func wsTicker(c *websocket.Conn, r *http.Request) {
	for {
		message := "tick: " + time.Now().String()

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		err := wsjson.Write(ctx, c, message)
		cancel()
		if err != nil {
			if websocket.CloseStatus(err) != websocket.StatusNormalClosure &&
				websocket.CloseStatus(err) != websocket.StatusGoingAway {
				log.Println(err)
			}
			return
		}

		time.Sleep(time.Second * 3)
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "")

	go wsTicker(c, r)
	wsReader(c, r)

	c.Close(websocket.StatusNormalClosure, "")
}

func MoveCounter(board *YacBoard.YacBoard, level int) int {
	var moves [256]YacBoard.Move
	moveCount := int(board.GetMoves(&moves))
	if level <= 1 {
		return moveCount
	}
	level--
	totalCount := 0
	bi := board.GetBoardInfo()
	for m := 0; m < moveCount; m++ {
		board.DoMove(moves[m])
		totalCount += MoveCounter(board, level)
		board.DoMoveBackward(moves[m], bi)
	}
	return totalCount
}

func testYacBoard() {
	var board YacBoard.YacBoard
	err := board.SetFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		panic(err)
	}

	for level := 1; level < 7; level++ {
		fmt.Print("Level: ", level, " Count: ")
		fmt.Println(MoveCounter(&board, level))
	}

	//var moves [256]YacBoard.Move
	//moveCount := board.GetMoves(&moves)

	//for i := 0; i < int(moveCount); i++ {
	//	fmt.Printf("    %3d - %v\n", i+1, moves[i])
	//}

	//fmt.Println(Crc64.FromBoard(&board))
}

func main() {
	//testYacBoard()
	//return

	_ = mime.AddExtensionType(".js", "application/javascript")
	//ct := mime.TypeByExtension(".js")
	//fmt.Printf("ct: %s\n", ct)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/ws", wsEndpoint)

	fmt.Println("run server: localhost:9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
