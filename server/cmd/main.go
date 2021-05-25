package main

import (
	"flag"

	"github.com/CristalT/gochat/server"
)

func main() {
	port := flag.String("port", "4000", "Server port")

	flag.Parse()

	s := server.NewServer()
	s.Listen(":" + *port)
	s.Start()
}
