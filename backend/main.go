//main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http server address")

func main() {
	flag.Parse()

	wsServer := NewWebsocketServer()
	go wsServer.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(wsServer, w, r)
	})

	fs := http.FileServer(http.Dir("./frontend/dist/"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 8080")
	log.Panic(
		http.ListenAndServe(*addr, nil),
	)
}
