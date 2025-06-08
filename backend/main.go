package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type Run struct {
	EventId     string `json:"eventId" db:"event_id"`
	RunNumber   int    `json:"runNumber,string" db:"run_number"`
	CarNumber   string `json:"carNumber" db:"car_number"`
	RawTime     string `json:"rawTime" db:"raw_time"`
	PaxTime     string `json:"paxTime" db:"pax_time"`
	CarClass    string `json:"carClass" db:"car_class"`
	DriverName  string `json:"driverName" db:"driver_name"`
	Cones       int    `json:"cones,string" db:"cones"`
	IsDnf       int    `json:"isDnf" db:"is_dnf"`
	GetsRerun   int    `json:"getsRerun" db:"gets_rerun"`
	LastUpdated int    `json:"lastUpdated" db:"last_updated"`
	Created     int    `json:"created" db:"created"`
}

type Event struct {
	EventId       string `json:"eventId" db:"event_id"`
	ClubName      string `json:"clubName" db:"club_name"`
	EventLocation string `json:"eventLocation" db:"event_location"`
	EventDate     string `json:"eventDate" db:"event_date"`
	EventNumber   int    `json:"eventNumber,string" db:"event_number"`
}

var db *sqlx.DB
var BASE_URL = os.Getenv("BASE_URL")

func getRuns_sql(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)
	EventId := params["EventId"]
	if EventId == "" {
		http.Error(w, "Missing event ID", http.StatusBadRequest)
		return
	}

	var newRuns []Run

	err := db.Select(&newRuns, "SELECT * FROM runs WHERE event_id = ? ORDER BY run_number DESC", EventId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Get runs error", http.StatusInternalServerError)
		return
	}

	if len(newRuns) == 0 {
		http.Error(w, "No runs found for event", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newRuns)
}

func getNewRuns_sql(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)
	LastUpdated, err := strconv.Atoi(params["LastUpdated"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Get last updated error", http.StatusBadRequest)
		return
	}

	EventId := params["EventId"]
	if EventId == "" {
		http.Error(w, "Missing event ID", http.StatusBadRequest)
		return
	}

	var newRuns []Run

	err = db.Select(&newRuns, "SELECT * FROM runs WHERE last_updated > ? AND event_id = ? ORDER BY run_number DESC", LastUpdated, EventId)
	if err != nil {
		log.Println("new runs error")
		log.Println(err)
		http.Error(w, "Get new runs error", http.StatusInternalServerError)
		return
	}

	if len(newRuns) == 0 {
		http.Error(w, "No new runs", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newRuns)
}

func getRun_sql(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)

	RunNumber, err := strconv.Atoi(params["RunNumber"])
	if err != nil {
		http.Error(w, "Get run error", http.StatusBadRequest)
		return
	}

	EventId := params["EventId"]
	if EventId == "" {
		http.Error(w, "Missing event ID", http.StatusBadRequest)
		return
	}

	var run Run

	err = db.Get(&run, "SELECT * FROM runs WHERE run_number = ? AND event_id = ?", RunNumber, EventId)
	if err != nil && err.Error() == "sql: no rows in result set" {
		http.Error(w, "No run found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Println(err)
		http.Error(w, "Get run error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(run)
}

func getCarsRuns_sql(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)

	CarNumber, err := strconv.Atoi(params["CarNumber"])
	if err != nil {
		http.Error(w, "Missing car number", http.StatusBadRequest)
		return
	}

	EventId := params["EventId"]
	if EventId == "" {
		http.Error(w, "Missing event ID", http.StatusBadRequest)
		return
	}

	var carRuns []Run

	err = db.Select(&carRuns, "SELECT * FROM runs WHERE car_number = ? AND event_id = ?", CarNumber, EventId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Get cars runs error", http.StatusInternalServerError)
		return
	}

	if len(carRuns) == 0 {
		http.Error(w, "Car runs not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carRuns)
}

func createRun_sql(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)

	EventId := params["EventId"]
	if EventId == "" {
		http.Error(w, "Missing event ID", http.StatusBadRequest)
		return
	}

	var run Run
	_ = json.NewDecoder(r.Body).Decode(&run)

	run.EventId = EventId

	log.Println(run)

	createTime := int(time.Now().Unix())

	run.LastUpdated = createTime
	run.Created = createTime

	_, err := db.NamedExec("INSERT INTO runs (event_id, run_number, car_number, raw_time, pax_time, car_class, driver_name, cones, is_dnf, gets_rerun, last_updated, created) VALUES (:event_id, :run_number, :car_number, :raw_time, :pax_time, :car_class, :driver_name, :cones, :is_dnf, :gets_rerun, :last_updated, :created)", run)
	if err != nil && strings.Contains(err.Error(), "Duplicate entry") {
		log.Println("Run already exists")
		http.Error(w, "Run already exists", http.StatusInternalServerError)
		return
	} else if err != nil {
		log.Println(err)
		http.Error(w, "Insert event error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(run)
}

func deleteRun_sql(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)
	RunNumber, err := strconv.Atoi(params["RunNumber"])
	if err != nil {
		http.Error(w, "Missing run number", http.StatusBadRequest)
		return
	}

	EventId := params["EventId"]
	if EventId == "" {
		http.Error(w, "Missing event ID", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM runs WHERE run_number = ? AND event_id = ?", RunNumber, EventId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Delete run error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Run deleted")
}

// todo: sql, and fully implement
func updateRun(w http.ResponseWriter, r *http.Request) {
	/*enableCors(&w)

	params := mux.Vars(r)

	var updatedRun Run
	_ = json.NewDecoder(r.Body).Decode(&updatedRun)

	log.Println(updatedRun)

	EventId := params["EventId"]
	if EventId == "" {
		http.Error(w, "Missing event ID", http.StatusBadRequest)
		return
	}

	if updatedRun.RunNumber == 0 {
		http.Error(w, "Run number is required", http.StatusBadRequest)
		return
	}

	for _, run := range runs {
		if run.RunNumber == updatedRun.RunNumber && run.EventId == EventId {
			updatedRun.LastUpdated = int(time.Now().Unix())
			updatedRun.Created = run.Created
			run = updatedRun
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedRun)*/
}

func getEvent_sql(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)

	EventId := params["EventId"]

	var event Event

	err := db.Get(&event, "SELECT * FROM events WHERE event_id = ?", EventId)
	if err != nil && err.Error() == "sql: no rows in result set" {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Println(err)
		http.Error(w, "Get event error", http.StatusInternalServerError)
		return
	}

	if event.EventId == EventId {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(event)
		return
	}

	http.Error(w, "Event not found", http.StatusNotFound)
}

func getEvents_sql(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var events_list []Event

	err := db.Select(&events_list, "SELECT * FROM events ORDER BY event_date DESC")
	if err != nil {
		log.Println(err)
		http.Error(w, "Get events error", http.StatusInternalServerError)
		return
	}

	if len(events_list) == 0 {
		http.Error(w, "No events found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events_list)
}

func createEvent_sql(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var event Event
	_ = json.NewDecoder(r.Body).Decode(&event)

	log.Println(event)

	if event.EventId == "random" {
		log.Println("Generating random event ID")
		u := uuid.New()
		event.EventId = u.String()
		log.Println(event)
	} else {
		log.Println("Using provided event ID")
	}

	_, err := db.NamedExec("INSERT INTO events (event_id, club_name, event_location, event_date, event_number) VALUES (:event_id, :club_name, :event_location, :event_date, :event_number)", event)

	if err != nil && strings.Contains(err.Error(), "Duplicate entry") {
		log.Println("Event already exists")
		http.Error(w, "Event already exists", http.StatusInternalServerError)
		return
	} else if err != nil {
		log.Println(err)
		http.Error(w, "Insert event error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func deleteEvent_sql(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)
	EventId := params["EventId"]
	if EventId == "" {
		http.Error(w, "Missing event ID", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("DELETE FROM events WHERE event_id = ?", EventId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Delete event error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Event deleted")
}

func getLeaderboardRuns_sql(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)
	EventId := params["EventId"]
	if EventId == "" {
		http.Error(w, "Missing event ID", http.StatusBadRequest)
		return
	}

	var leaderboardRuns []Run

	err := db.Select(&leaderboardRuns, "SELECT * FROM runs WHERE pax_time > 0 AND event_id = ? AND is_dnf == 0 AND gets_rerun == 0", EventId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Insert event error", http.StatusInternalServerError)
		return
	}

	if len(leaderboardRuns) == 0 {
		http.Error(w, "Leaderboard runs not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leaderboardRuns)
}

func getEventClasses_sql(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)
	EventId := params["EventId"]
	if EventId == "" {
		http.Error(w, "Missing event ID", http.StatusBadRequest)
		return
	}

	var classes []string

	err := db.Select(&classes, "SELECT DISTINCT car_class FROM runs WHERE event_id = ? AND car_class IS NOT NULL AND car_class != '' ORDER BY car_class", EventId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Get event classes error", http.StatusInternalServerError)
		return
	}

	if len(classes) == 0 {
		http.Error(w, "No classes found for event", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classes)
}

func getClass_sql(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)
	EventId := params["EventId"]
	ClassName := params["ClassName"]

	if EventId == "" || ClassName == "" {
		http.Error(w, "Missing event ID or class name", http.StatusBadRequest)
		return
	}

	var classRuns []Run

	err := db.Select(&classRuns, "SELECT * FROM runs WHERE event_id = ? AND car_class = ? ORDER BY pax_time ASC", EventId, ClassName)
	if err != nil {
		log.Println(err)
		http.Error(w, "Get class runs error", http.StatusInternalServerError)
		return
	}

	if len(classRuns) == 0 {
		http.Error(w, "No runs found for class", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classRuns)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	http.Error(w, "404 page not found", http.StatusNotFound)
}

func optionsCors(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", BASE_URL)
}

func main() {

	DB_PASSWORD, err := os.ReadFile("/run/secrets/DB_PASSWORD")
	if err != nil {
		log.Fatalln("DB password not found")
	}
	CONNECTION_STRING := fmt.Sprintf("autocross:%s@tcp(db:3306)/autocross", DB_PASSWORD)

	retries := 0

	for retries < 10 {
		db, err = sqlx.Connect("mysql", CONNECTION_STRING)
		if err == nil {
			break
		} else if retries == 4 {
			log.Fatalln(err)
		} else {
			retries++
			time.Sleep(5 * time.Second)
		}
	}

	log.Println("Connected to database")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/api/runs/{EventId}/", getRuns_sql).Methods("GET")
	r.HandleFunc("/api/runs/{EventId}/{LastUpdated}", getNewRuns_sql).Methods("GET")
	r.HandleFunc("/api/run/{EventId}/{RunNumber}", getRun_sql).Methods("GET")
	r.HandleFunc("/api/runs/{EventId}", createRun_sql).Methods("POST")
	r.HandleFunc("/api/runs/{EventId}", optionsCors).Methods("OPTIONS")
	r.HandleFunc("/api/runs/{EventId}/{RunNumber}", updateRun).Methods("POST")
	r.HandleFunc("/api/runs/{EventId}/{RunNumber}", optionsCors).Methods("OPTIONS")
	r.HandleFunc("/api/car/{EventId}/{CarNumber}", getCarsRuns_sql).Methods("GET")
	r.HandleFunc("/api/runs/leaderboard/{EventId}/", getLeaderboardRuns_sql).Methods("GET")
	r.HandleFunc("/api/event/{EventId}", getEvent_sql).Methods("GET")
	r.HandleFunc("/api/class/{EventId}", getEventClasses_sql).Methods("GET")
	r.HandleFunc("/api/class/{EventId}/{ClassName}", getClass_sql).Methods("GET")
	r.HandleFunc("/api/events/", getEvents_sql).Methods("GET")
	r.HandleFunc("/api/events", createEvent_sql).Methods("POST")
	r.HandleFunc("/api/events", optionsCors).Methods("OPTIONS")
	r.HandleFunc("/api/event/delete/{EventId}", deleteEvent_sql).Methods("GET")
	r.HandleFunc("/api/run/delete/{EventId}/{RunNumber}", deleteRun_sql).Methods("GET")

	fmt.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
