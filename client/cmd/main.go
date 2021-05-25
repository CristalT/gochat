package main

import (
	"flag"
	"log"

	"github.com/CristalT/gochat/client"
	"github.com/CristalT/gochat/tui"
)

func main() {
	address := flag.String("server", "", "Server address")

	flag.Parse()

	client := client.NewClient()
	err := client.Dial(*address)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()
	go client.Start()

	tui.StartUi(client)
}
