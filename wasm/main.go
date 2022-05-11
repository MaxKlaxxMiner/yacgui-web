package main

import (
	"context"
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/YacBoard"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"syscall/js"
	"time"
	. "wasm/WasmGlobal"
)

func loglog(line string) {
	js.Global().Get("wg").Call("loglog", line)
}

func wsReader(c *websocket.Conn) {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		var v any
		err := wsjson.Read(ctx, c, &v)
		cancel()
		if err != nil {
			if websocket.CloseStatus(err) != websocket.StatusNormalClosure &&
				websocket.CloseStatus(err) != websocket.StatusGoingAway {
				loglog(err.Error())
			}
			return
		}

		loglog("received: " + fmt.Sprint(v))
	}
}

func wsPinger(c *websocket.Conn) {
	var board YacBoard.YacBoard
	err := board.SetFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		panic(err)
	}

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

		err := wsjson.Write(ctx, c, "fen: "+board.GetFEN())
		cancel()
		if err != nil {
			break
		}

		time.Sleep(time.Second * 3)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	c, _, err := websocket.Dial(ctx, "ws://localhost:9090/ws", nil)
	cancel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "")

	WgInit()

	go wsReader(c)
	wsPinger(c)

	c.Close(websocket.StatusNormalClosure, "")

	//c := make(chan struct{}, 0)
	//<-c
}
