package seadisc

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var (
	BaseUrl string = "wss://{network}stream.openseabeta.com/socket/websocket?token="
)

type Sea struct {
	SocketUrl string
}

func (s *Sea) SetSocketUrl(network, token string) {
	var net string = ""

	if strings.ToLower(network) != "main" {
		net = "testnets-"
	}

	s.SocketUrl = strings.Replace(BaseUrl, "{network}", net, 1)
}

func (s *Sea) OpenseaSocket() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Connect to websocket
	c, _, err := websocket.DefaultDialer.Dial(s.SocketUrl, nil)
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
			// log.Printf("recv: %s", message)
			Discord(message)
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
