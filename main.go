package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

type Metal struct {
	Name string
}

type Experiment struct {
	Metal string
	ExpectedCycles int
	Data [][]int
	Failed bool
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	mux.HandleFunc("/metals", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method is not allowed", 405)
			return
		}

		var d []Metal
		d = append(d, Metal{Name: "Aluminum"})
		d = append(d, Metal{Name: "Copper"})
		d = append(d, Metal{Name: "Silver"})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(d)
	})

	mux.HandleFunc("/new-experiment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", 405)
			return
		}

		type Body struct {
			Metal string `json:"metal"`
			Cycles int `json:"cycles"`
		}

		var b Body
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Not able to decode json")
			http.Error(w, "Error decoding json", http.StatusBadRequest)
			return
		}

		experiment := genExperiment(b.Metal, b.Cycles)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(experiment)

		fmt.Println(experiment)


	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "4321"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}


	fmt.Printf("Starting server on port: %v\n", port)

	srv.ListenAndServe()
}


func genExperiment(material string, cycleCt int) Experiment {
	var exp Experiment
	exp.Metal = material
	exp.ExpectedCycles = cycleCt


	for i:=0; i < cycleCt; i++ {
		if exp.Failed {
			break
		}

		var run []int
		run = append(run, 700)
	
		for j:=1; j < 10; j++ {
			newT := run[j-1]
			ran := rand.Intn(100)
			if ran < 50 && newT <= 750 {
				newT +=10
			}

			if ran >=50 && ran <=90 && newT > 690 && newT < 750 {
				newT -= 3
			}
			
			if newT > 750 {
				exp.Failed = true
			}

			run = append(run, newT)
		}

		exp.Data = append(exp.Data, run)
	}
	return exp
}
