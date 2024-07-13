package main

import "fmt"

// The select statement listens on multiple channels
//  and performs actions based on which channel receives a message.

func main() {

	channel := make(chan string, 3) //buffered channel

	characters := []string{"a", "b", "c"}

	for _, s := range characters {
		select {
		case channel <- s:
		}
	}

	close(channel)

	for result := range channel {
		fmt.Println(result)
	}
}

// ==================== EXAMPLE OF A CHANNEL WITH MULTIPLE CASES ====================

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}
		}
	}
}
