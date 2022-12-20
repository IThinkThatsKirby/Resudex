package handler

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func React(w http.ResponseWriter, r *http.Request) {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write
			err := websocket.Message.Send(ws, "document.getElementById('1').style.backgroundColor = 'red'")
			if err != nil {
				println(err.Error())
				ws.Close()
			}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				println(err.Error())
			}
			fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(w, r)
}
