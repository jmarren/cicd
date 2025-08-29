package internal

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8090"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi there"))
}

func Run() {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(rootHandler))

	fmt.Printf("listening on %s...\n", PORT)
	if err := http.ListenAndServe(PORT, mux); err != nil {
		log.Fatalf("Failed to listen with error: %s", err)
	}
}
