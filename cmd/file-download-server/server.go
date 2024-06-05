package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func handleFileDownload(dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract file name from URL
		fileName := strings.TrimPrefix(r.URL.Path, "/download-")
		log.Println("Download ", fileName)
		fileName = filepath.Clean(fileName) // Clean the file path

		// Create the full file path
		filePath := filepath.Join(dir, fileName)

		// Open the file
		file, err := os.Open(filePath)
		if err != nil {
			log.Println("ERR: File not found")
			http.Error(w, "File not found.", http.StatusNotFound)
			return
		}
		defer file.Close()

		// Get the file info
		fileInfo, err := file.Stat()
		if err != nil {
			http.Error(w, "Could not get file info.", http.StatusInternalServerError)
			return
		}

		// Set the headers
		w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

		// Stream the file content to the response
		http.ServeContent(w, r, fileName, fileInfo.ModTime(), file)
	}
}

type Server struct {
	server *http.Server
}

func (s *Server) Start(host string, port string, dir string) {
	address := fmt.Sprintf("%s:%s", host, port)
	log.Printf("File directory:%s\n", dir)
	log.Printf("Starting server on %s\n", address)

	handler := &RegexpHandler{}
	reg1, _ := regexp.Compile(`/download-([0-9]+)(MB|GB|mb|gb)(.bin)*`)
	handler.HandleFunc(reg1, handleFileDownload(dir))

	s.server = &http.Server{
		Addr:    address,
		Handler: handler,
	}
	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalf("server err: %s\n", err.Error())
	}

}
func (s *Server) Shutdown(ctx context.Context) error {
	// Shutdown the server gracefully
	return s.server.Shutdown(ctx)
}
