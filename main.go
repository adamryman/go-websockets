package main

import (
	"golang.org/x/net/websocket"
	"net/http"
	"time"
)

var msg = make(chan []byte)

func wsHandler(ws *websocket.Conn) {
	for {
		wsMsg := <-msg
		ws.Write(wsMsg)
	}
}

func buttonPressed(w http.ResponseWriter, r *http.Request) {
	for {
		time.Sleep(time.Second)
		msg <- []byte("hello, world")
	}
}

func main() {
	http.HandleFunc("/button", buttonPressed)
	http.Handle("/socket", websocket.Handler(wsHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
