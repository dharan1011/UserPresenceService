package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func main() {
	var context = context.Background()
	var redisClient = CreateRedisClient(context, 100)
	// defer redisClient.db.Close()
	// defer log.Println("Worked")
	var service = CreateUserPresenceService(redisClient, time.Duration(30)*time.Second)
	r := mux.NewRouter()
	r.HandleFunc("/api/healthCheck", logging(HealthCheckHandler)).Methods("GET")
	r.HandleFunc("/api/notifyUserPresence", logging(service.HandlerNotifyUserPresence)).Methods("POST")
	r.HandleFunc("/api/getUserPresence", logging(service.HandlerGetUserPresence)).Methods("GET")
	httpServer := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}
	log.Println("Server listening on port", 8080)
	httpServer.ListenAndServe()
}

func HealthCheckHandler(rw http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "OK",
	}
	json.NewEncoder(rw).Encode(response)
}
