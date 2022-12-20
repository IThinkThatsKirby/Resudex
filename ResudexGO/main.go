package main

import (
	"net/http"
	ws "resudex/api/handler"
)

func main() {
	// serve html file for site entry
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/ws", ws.React)

	server := http.ListenAndServe(":1234", nil)
	println("Listening on port 1234" + server.Error())
}
