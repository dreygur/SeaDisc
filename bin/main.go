package main

import (
	"os"

	"github.com/dreygur/seadisc"
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

	// Start the Websocket
	var socket seadisc.Sea
	socket.SetSocketUrl(network, token)
	socket.OpenseaSocket()
}
