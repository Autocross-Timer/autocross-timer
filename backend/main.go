package main

import (
    "encoding/json"
	"fmt"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

type Run struct {
	RunNumber int `json:"runNumber"`
	CarNumber int `json:"carNumber"`
	RawTime  int `json:"rawTime"`
	PaxTime  int `json:"paxTime"`
	Class   string `json:"class"`
	Name   string `json:"name"`
	Cones   int `json:"cones"`
}

var runs []Run

func getRuns(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if len(runs) == 0 {
		http.Error(w, "No runs", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(runs)
}

func getNewRuns(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)
	LastRunNumber, err := strconv.Atoi(params["LastRunNumber"])
	if err != nil {
		http.Error(w, "Get new runs error", http.StatusBadRequest)
		return
	}

	var newRuns []Run

	for _, run := range runs {
		if run.RunNumber > LastRunNumber {
			newRuns = append(newRuns, run)
		}
	}

	if len(newRuns) == 0 {
		http.Error(w, "No new runs", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newRuns)
}

func getRun(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)

	RunNumber, err := strconv.Atoi(params["RunNumber"])
	if err != nil {
		http.Error(w, "Get run error", http.StatusBadRequest)
		return
	}

	for _, item := range runs {
		if item.RunNumber == RunNumber {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Run not found", http.StatusNotFound)
}

func getCarsRuns(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)

	CarNumber, err := strconv.Atoi(params["CarNumber"])
	if err != nil {
		http.Error(w, "Get car runs error", http.StatusBadRequest)
		return
	}

	var carRuns []Run

	for _, run := range runs {
		if run.CarNumber == CarNumber {
			carRuns = append(carRuns, run)
		}
	}

	if len(carRuns) == 0 {
		http.Error(w, "Car runs not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carRuns)
}

func createRun(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var run Run
	_ = json.NewDecoder(r.Body).Decode(&run)

	log.Println(run)

	if run.RunNumber == 0 {
		http.Error(w, "Run number is required", http.StatusBadRequest)
		return
	}

	runs = append(runs, run)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(run)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {

	runs = append(runs, Run{RunNumber: 1, CarNumber: 1, RawTime: 100, PaxTime: 90, Class: "A Street", Name: "John Doe", Cones: 0})
	runs = append(runs, Run{RunNumber: 2, CarNumber: 16, RawTime: 98, PaxTime: 85, Class: "B Street", Name: "Jared Kelly", Cones: 2})

	r := mux.NewRouter()

	r.HandleFunc("/runs/", getRuns).Methods("GET")
	r.HandleFunc("/runs/{LastRunNumber}", getNewRuns).Methods("GET")
	r.HandleFunc("/run/{RunNumber}", getRun).Methods("GET")
	r.HandleFunc("/runs", createRun).Methods("POST")
	r.HandleFunc("/car/{CarNumber}", getCarsRuns).Methods("GET")

	fmt.Println("Server is running on port 8000...")
    log.Fatal(http.ListenAndServe(":8000", r))
}