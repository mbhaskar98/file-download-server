package main

import (
	"flag"
)

var (
	host string
	port string
	dir  string
)

func init() {
	flag.StringVar(&host, "host", "localhost", "Hostname to serve on")
	flag.StringVar(&port, "port", "8080", "Port to serve on")
	flag.StringVar(&dir, "dir", ".", "Directory containing the files to serve")
}

func main() {
	flag.Parse()

	s := &Server{}
	s.Start(host, port, dir)
}
