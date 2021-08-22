package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const PORT = "8000"

func GetDateNow(w http.ResponseWriter, r *http.Request) {
	var currentTime = time.Now()
	var response = make(map[string]string)

	response["date-now"] = currentTime.String()

	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/date-now", GetDateNow).Methods("GET")

	fmt.Println("Starting Restful services...")
	fmt.Println("Using port:" + PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
