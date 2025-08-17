package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Tracker struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Status string `json:"status"`
}

type TrackerCollection struct {
	Trackers []Tracker `json:"trackers"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var trackers = TrackerCollection{
	Trackers: []Tracker{
		{ID: "1", Name: "Tracker 1", Status: "Online"},
		{ID: "2", Name: "Tracker 2", Status: "Offline"},
		{ID: "3", Name: "Tracker 3", Status: "Online"},
	},
}

func main() {
	http.HandleFunc("/trackers", handleTrackers)
	http.ListenAndServe(":8080", nil)
}

func handleTrackers(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	for {
		select {
		case <-time.After(5 * time.Second):
			jsonTrackers, err := json.Marshal(trackers)
			if err != nil {
				log.Println(err)
				return
			}
			ws.WriteMessage(websocket.TextMessage, jsonTrackers)
		}
	}
}