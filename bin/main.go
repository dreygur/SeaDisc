package main

import (
	"fmt"
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

	// Get the Socket URL
	url := seadisc.GetSocketUrl(network, token)
	fmt.Println(url)
}
