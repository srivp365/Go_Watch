package main

import (
	"fmt"
	"sync"
	"time"
)

type Hub struct {
	rooms    map[string]*Room
	eventCh  chan RoomEvent
	tickerCh <-chan time.Time
	mu       sync.Mutex
}

func (r *Room) broadcastToAll (msg []byte, excludeClientID string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for id, client := range r.clientIDs {
		if id == excludeClientID {
			continue
		}
		
		err := 
		
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
			case "join":
				room.clientIDs[]
			case "seek":
				room.timestamp = event.timestamp

			}
		case heartbeat := <- h.tickerCh:
			// something
		}
	}
}

type RoomEvent struct {
	roomID    string
	videoID   string
	timestamp float64
	event     string
	clientIDs map[string]string
}

type Room struct {
	roomID    string
	clientIDs map[string]string
	videoID   string
	timestamp float64
	playing   bool
	mu        sync.Mutex
}

func (r *Room) setPlaying(playing bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.playing = playing
}

func (r *Room) getState() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.playing
}

func main() {
	fmt.Println("Hello, World")
}
