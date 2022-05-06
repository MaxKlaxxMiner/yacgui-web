package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"mime"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsReader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))

		p = []byte("löl: " + string(p))

		if err = conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsTicker(conn *websocket.Conn) {
	for {
		message := "tick: " + time.Now().String()
		if err := conn.WriteMessage(1, []byte(message)); err != nil {
			return
		}
		time.Sleep(time.Second * 3)
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")

	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}

	go wsTicker(ws)
	wsReader(ws)
}

func main() {
	_ = mime.AddExtensionType(".js", "application/javascript")
	//ct := mime.TypeByExtension(".js")
	//fmt.Printf("ct: %s\n", ct)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/ws", wsEndpoint)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
