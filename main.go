package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Metal struct {
	Name string
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	mux.HandleFunc("/metals", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method is not allowed", 405)
		}

		var d []Metal
		d = append(d, Metal{Name: "Aluminum"})
		d = append(d, Metal{Name: "Copper"})
		d = append(d, Metal{Name: "Silver"})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(d)
	})


	port := os.Getenv("PORT")
	if port == "" {
		port = "4321"
	}

	srv := &http.Server{
		Addr: ":"+port,
		Handler: mux,
	}

	fmt.Printf("Starting server on port: %v", port)

	srv.ListenAndServe()
}
