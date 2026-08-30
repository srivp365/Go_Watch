package main

import (
	"sync"
	"time"
)

type Hub struct {
	roomIDs map[string]*Room
	mu      sync.Mutex
}

type Room struct {
	roomID    string
	clientIDs map[string]*Client
	videoID   string
	timestamp float64
	playing   bool
	eventCh   chan *RoomEvent
	tickerCh  <-chan time.Time
	mu        sync.Mutex
}

type RoomEvent struct {
	roomID    string
	videoID   string
	timestamp float64
	event     string
	client    *Client
}

// create / check room => ccRoom
func (h *Hub) ccRoom(roomID string, client *Client) (r *Room) {
	h.mu.Lock()

	room, ok := h.roomIDs[roomID]

	if !ok {
		room = createRoom(roomID)
		h.roomIDs[roomID] = room
		go room.Run()

	}

	re := &RoomEvent{
		roomID: roomID,
		event:  "join",
		client: client,
	}

	room.eventCh <- re

	h.mu.Unlock()

	return
}

func createRoom(roomID string) (r *Room) {
	newRoom := &Room{
		roomID:    roomID,
		clientIDs: make(map[string]*Client),
		videoID:   "",
		timestamp: 0,
		playing:   false,
		eventCh:   make(chan *RoomEvent),
		tickerCh:  make(<-chan time.Time),
	}

	return newRoom
}

func (r *Room) setPlaying(playing bool) {
	// ensure only 1 client can pause / play at once
	r.mu.Lock()
	defer r.mu.Unlock()
	r.playing = playing
}

// func (r *Room) broadcastToAll(msg []byte, excludeClientID string) {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()

// 	for id, client := range r.clientIDs {
// 		if id == excludeClientID {
// 			continue
// 		}

// 		err := client.WriteMessage(websocket.TextMessage, msg)
// 		if err != nil {
// 			log.Printf("broadcast failed for user %s: %v", id, err)
// 		}

// 	}

// }

func (r *Room) Run() {
	for {
		select {
		case event := <-r.eventCh:
			switch event.event {
			case "play":
				r.setPlaying(true)
			case "pause":
				r.setPlaying(false)
			case "leave":
				delete(r.clientIDs, event.client.id)
			case "join":
				r.clientIDs[event.client.id] = event.client
			case "seek":
				r.timestamp = event.timestamp
			}
		case heartbeat := <-r.tickerCh:
			// i know this is incorrect, will fix it.
			r.timestamp = float64(heartbeat.Second())
		}
	}
}
