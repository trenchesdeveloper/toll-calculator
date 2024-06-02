package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/trenchesdeveloper/toll-calculator/types"
)

const sendInterval = time.Second * 5

const wsEndpont = "ws://localhost:8080/ws"



func generateCoordinates() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}

func generateLocation() (float64, float64) {
	return generateCoordinates(), generateCoordinates()
}

func generateOBUIDS(n int) []int {
	var ids []int
	for i := 0; i < n; i++ {
		ids = append(ids, rand.Intn(1000))
	}
	return ids
}

func sendOBUData(conn *websocket.Conn, data types.OBUData) {
	conn.WriteJSON(data)
}

func main() {
	obuIds := generateOBUIDS(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpont, nil)

	if err != nil {
		log.Fatal("Error connecting to server: ", err)
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		for _, id := range obuIds {
			lat, long := generateLocation()
			data := types.OBUData{OBUID: id, Lat: lat, Long: long}
			fmt.Println(data)

			if err := conn.WriteJSON(data); err != nil {
				log.Fatal("Error sending data: ", err)
			}
		}
		fmt.Println("Sending data...", generateCoordinates())
		time.Sleep(sendInterval)
	}
}
