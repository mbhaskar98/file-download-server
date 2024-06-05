package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	go func() {

		s.Start(host, port, dir)
	}()

	// Register interrupt signal handler
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	// Wait for interrupt signal
	<-stop

	// Shutdown the server gracefully
	log.Println("Shutting down server...")
	if err := s.Shutdown(context.TODO()); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server stopped.")
}
