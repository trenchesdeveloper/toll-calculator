package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/trenchesdeveloper/toll-calculator/types"
)

type DataReceiver struct {
	msgch chan types.OBUData
	conn *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin returns true if the request Origin header is acceptable.
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (dr *DataReceiver) echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // Upgrade the HTTP server connection to the WebSocket protocol.
	if err != nil {
		log.Println(err)
		return
	}

	dr.conn = conn

	go dr.wsReceiveLoop()
}

func (dr *DataReceiver) wsReceiveLoop() {
	log.Println("New OBU connected")
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println(err)
			continue
		}
		log.Printf("Received data from: %d\n", data.OBUID)
		dr.msgch <- data
	}
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgch: make(chan types.OBUData, 128),
	}
}

func main() {
	dr := NewDataReceiver()
	http.HandleFunc("/ws", dr.echoHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
