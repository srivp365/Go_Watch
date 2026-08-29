package main

import (
	"fmt"
	"time"
)

type Hub struct {
	rooms    map[string]*Room
	eventCh  chan RoomEvent
	tickerCh <-chan time.Time
}

func newHub() *Hub {
	return &Hub{
		rooms:    make(map[string]*Room),
		eventCh:  make(chan RoomEvent),
		tickerCh: make(<-chan time.Time),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case event := <-h.eventCh:
			room, ok := h.rooms[event.roomID]
			if !ok {
				fmt.Println("* I'm Batman *")
				continue
			}
			switch event.event {
			case "play":
				room.setPlaying(true)
			case "pause":
				room.setPlaying(false)
			case "leave":
				delete(room.clientIDs, event.clientId)
			case "join":
				room.clientIDs[event.clientId] = event.conn
			case "seek":
				room.timestamp = event.timestamp

			}
		case heartbeat := <-h.tickerCh:
			// something
		}
	}
}
