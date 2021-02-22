package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func helloPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func reader(ws *websocket.Conn) {
	for {
		mt, p, err := ws.ReadMessage()
		if err != nil {
			log.Print(err)
			return
		}
		fmt.Println(string(p))

		if err = ws.WriteMessage(mt, p); err != nil {
			log.Print(err)
			return
		}
	}
}

func webSock(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "hello world")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
	}

	fmt.Println("Client Successfully Connected....")

	reader(wsConn)
}

func setupRoutes() {
	http.HandleFunc("/", helloPage)
	http.HandleFunc("/ws", webSock)
}

func main() {
	fmt.Println("Testing Go WebSockets")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
