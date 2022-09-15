package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"web-proxy/websocketporxy"
)

func main() {
	parse, _ := url.Parse("http://debian:8123")
	proxy := websocketporxy.NewProxy(parse)
	proxy.ReceiveHandle = func(messageType int, msg []byte, con *websocket.Conn) {
		con.WriteMessage(messageType, msg)
		fmt.Println(string(msg))
	}
	http.Handle("/", proxy)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
