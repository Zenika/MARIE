package main

import (
	"log"

	"net/http"

	"github.com/Zenika/MARIE/backend/network"
)

func main() {
	http.HandleFunc("/ws", network.Handle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
