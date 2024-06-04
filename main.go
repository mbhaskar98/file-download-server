package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"``
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Usage: %s %s %s %s", os.Args[0], "HOST", "PORT", "FILEPATH")
	}
	// host := "192.168.1.2"
	host := os.Args[1]                  //"192.168.1.3"
	port, _ := strconv.Atoi(os.Args[2]) // 8000
	filePath := os.Args[3]
	address := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("Starting the server on: {%s}\n", address)
	http.HandleFunc("/", sendFile(filePath))
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("Error while starting server:", err)
	}

}

func sendFile(path string) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		if len(path) == 0 {
			path = "~/Desktop/10gb.bin"
		}

		tempArray := strings.Split(path, "/")
		filename := tempArray[len(tempArray)-1]

		fmt.Printf("Sending the file:%s to %s\n", filename, r.RemoteAddr)
		w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename))
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeFile(w, r, path)
	}
}

// func sendFile(w http.ResponseWriter, r *http.Request) {
// 	filename := "10gb.bin"
// 	temp := "/Users/bhaskarmahajan/Desktop/Tasks/HelperPrograms/file-download-server/"
// 	filePath := fmt.Sprintf("%s/%s", temp, filename)
// 	fmt.Printf("Sending the file:%s to %s\n", filename, r.RemoteAddr)
// 	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename))
// 	w.Header().Set("Content-Type", "application/octet-stream")
// 	http.ServeFile(w, r, filePath)
// }
