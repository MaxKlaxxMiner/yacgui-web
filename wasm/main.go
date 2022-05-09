package main

import (
	"context"
	"fmt"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"syscall/js"
	"time"
)

func loglog(line string) {
	js.Global().Call("loglog", line)
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
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

		err := wsjson.Write(ctx, c, "ping")
		cancel()
		if err != nil {
			break
		}

		time.Sleep(time.Second * 3)
	}
}

func main() {
	js.Global().Set("GoKeyDown", js.FuncOf(GoKeyDown))

	//js.Global().Set("balla", js.FuncOf(balla))
	//// balla = function() { this. }
	//

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	c, _, err := websocket.Dial(ctx, "ws://localhost:9090/ws", nil)
	cancel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "")

	go wsReader(c)
	wsPinger(c)

	c.Close(websocket.StatusNormalClosure, "")

	////js.Global().Get("console").Call("log", "muh")

	//c := make(chan struct{}, 0)
	//<-c
}

func GoKeyDown(_ js.Value, args []js.Value) any {
	if len(args) != 2 || args[0].Type() != js.TypeString || args[1].Type() != js.TypeBoolean {
		return false
	}
	code := args[0].String()
	defaultPrevented := args[1].Bool()

	if defaultPrevented {
		return false
	}

	if code == "KeyF" {
		js.Global().Get("ground").Call("toggleOrientation")
		return true
	}

	return false
}
