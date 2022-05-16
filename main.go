package main

import (
	"context"
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/YacBoard"
	"log"
	"mime"
	"net"
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

		time.Sleep(time.Second * 8)
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

func SendLine(client *net.TCPConn, line string) {
	buf := make([]byte, len(line)+2)
	buf[0] = byte(len(line))
	buf[1] = byte(len(line) >> 8)
	copy(buf[2:], line)
	client.Write(buf)
}

func testTcp() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:50354")
	if err != nil {
		panic(err)
	}
	client, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err)
	}
	client.SetNoDelay(true)

	go func() {
		buf := make([]byte, 2048)
		collect := make([]byte, 0)
		for {
			read, err := client.Read(buf)
			if err != nil {
				return
			}
			collect = append(collect, buf[0:read]...)
			for len(collect) > 1 {
				bytes := int(collect[0]) + int(collect[1])*256
				if bytes+2 > len(collect) {
					break
				}
				line := collect[2 : bytes+2]
				fmt.Println(string(line))
				collect = collect[bytes+2:]
			}
		}
	}()

	for {
		var line string
		fmt.Scanln(&line)
		if line == "quit" {
			break
		}
		SendLine(client, line)
	}
	client.Close()
	time.Sleep(time.Second)
}

func main() {
	//testYacBoard()
	//return

	testTcp()
	return

	_ = mime.AddExtensionType(".js", "application/javascript")

	http.Handle("/", http.FileServer(http.Dir("http-content/")))
	http.HandleFunc("/ws", wsEndpoint)

	fmt.Println("run server: localhost:9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
