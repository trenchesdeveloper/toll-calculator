package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/trenchesdeveloper/toll-calculator/types"
)

var kafkaTopic = "obudata"

type DataReceiver struct {
	msgch chan types.OBUData
	conn  *websocket.Conn
	prod  DataProducer
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

		if err := dr.produceData(data); err != nil {
			log.Println(err)
			continue
		}
	}
}

func NewDataReceiver() (*DataReceiver, error) {
	var( p DataProducer
		err error
		kafkaTopic = "obudata"
	)

	p, err = NewKafkaProducer(kafkaTopic)
	if err != nil {
		return nil, err
	}

	p = NewLogMiddleware(p)

	return &DataReceiver{
		msgch: make(chan types.OBUData, 128),
		prod:  p,
	}, nil
}

func (dr *DataReceiver) produceData(data types.OBUData) error {
	return dr.prod.ProduceData(data)
}

func main() {
	recv, err := NewDataReceiver()

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/ws", recv.echoHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
