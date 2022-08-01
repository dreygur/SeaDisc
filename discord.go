package seadisc

import (
	"encoding/json"
	"log"
)

type SeaEvent struct {
	Event   string `json:"event"`
	Payload struct {
		EventType string `json:"item_transferred"`
		Payload   struct {
			Collection struct {
				Slug string `json:"slug"`
			} `json:"collection"`
			EventTimestamp string `json:"event_timestamp"`
			FromAccount    struct {
				Address string `json:"address"`
			} `json:"from_account"`
			ToAccount struct {
				Address string `json:"address"`
			} `json:"to_account"`
			Item struct {
				Chain struct {
					Name string `json:"name"`
				} `json:"chain"`
				MetaData struct {
					AnimationUrl string `json:"animation_url"`
					ImageUrl     string `json:"image_url"`
					MetaDataUrl  string `json:"metadata_url"`
					Name         string `json:"name"`
				} `json:"metadata"`
				NFTid     string `json:"nft_id"`
				Permalink string `json:"permalink"`
			} `json:"item"`
			Quantity    int `json:"quantity"`
			Transaction struct {
				Hash      string `json:"hash"`
				Timestamp string `json:"timestamp"`
			} `json:"transaction"`
			SentAt string `json:"sent_at"`
		} `json:"payload"`
		Ref   string `jspn:"ref"`
		Topic string `json:"topic"`
	} `json:"payload"`
}

func Discord(message []byte) {
	var event SeaEvent
	_ = json.Unmarshal(message, &event)
	log.Printf("%v : %v : %v", event.Event, event.Payload.Payload.FromAccount.Address, event.Payload.Payload.ToAccount.Address)
	// log.Printf("recv: %s", message)
}
