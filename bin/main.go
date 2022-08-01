package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/dreygur/seadisc"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

func init() {
	// Load the .env file I have in project directory
	godotenv.Load()
}

func main() {
	// Initialize some Environment Variables
	var (
		network string = os.Getenv("OPENSEA_NETWORK")
		token   string = os.Getenv("OPENSEA_API_KEY")
	)

	// Get the Socket URL
	url := seadisc.GetSocketUrl(network, token)
	fmt.Println(url)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Connect to websocket
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	// Heartbit
	heartbitSchema := map[string]interface{}{
		"topic":   "phoenix",
		"event":   "heartbeat",
		"payload": map[string]interface{}{},
		"ref":     0,
	}
	heartbit, _ := json.Marshal(heartbitSchema)

	// Subscribe to all event
	eventSchema := map[string]interface{}{
		"topic":   "collection:*",
		"event":   "phx_join",
		"payload": map[string]interface{}{},
		"ref":     0,
	}
	event, _ := json.Marshal(eventSchema)
	_ = c.WriteMessage(websocket.TextMessage, event)

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, heartbit)
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
