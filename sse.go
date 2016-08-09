package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ServerSentEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return

	}
	msg := "anuchit"
	now := time.Now()
	fmt.Fprintf(w, "data: Message: %s time: %v\n\n", msg, now)
	log.Printf("Sent message %v ", now)
	time.Sleep(5 * time.Millisecond)
	f.Flush()
}

func main() {
	http.HandleFunc("/sse", ServerSentEventHandler)
	log.Println("started")
	log.Fatal(http.ListenAndServe(":8765", nil))
}
