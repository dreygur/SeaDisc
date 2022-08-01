package seadisc

import "log"

func Discord(message []byte) {
	log.Printf("recv: %s", message)
}
