package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/events", SSE)
	http.ListenAndServe(":8080", nil)
}

func SSE(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "SSE not supported!", http.StatusInternalServerError)
		return
	}
	defer r.Context().Done()
	defer fmt.Fprint(w, "data: end\n\n")
	defer flusher.Flush()

	log.Println("request!")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	timeout := time.After(time.Second * 1)
	for {
		select {
		case <-timeout:
			log.Println("timeout!")
			return
		default:
			fmt.Fprint(w, "data: ฅ^•ﻌ•^ฅ\n\n")
			flusher.Flush()
		}
	}
}
