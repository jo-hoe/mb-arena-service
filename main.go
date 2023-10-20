package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gorilla/mux"
	"github.com/jo-hoe/mb-arena-service/app"
)

var cache = make([]app.MBEvent, 0)

func main() {
	// schedule periodic updates
	scheduleUpdates()

	serve()
}

func scheduleUpdates() {
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Cron(getEnv("CACHE_UPDATE_CRON", "0 2 * * *")).Do(updateItems)
	if err != nil {
		log.Printf("could not schedule update job, error: %+v", err)
	}

	scheduler.StartAsync()
}

func updateItems() {
	log.Print("updating events")
	items, err := app.Spider(http.DefaultClient)
	if err != nil {
		log.Printf("could not spider events, error: %+v", err)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Start.Unix() < items[j].Start.Unix()
	})

	if len(items) == 0 {
		log.Print("no events to add")
	} else {
		cache = items
		log.Printf("added %d events to cache", len(items))
	}
}

func createAPIConfig() (router *mux.Router) {
	// config API
	router = mux.NewRouter()
	router.HandleFunc("/", DefaultHandler)
	http.Handle("/", router)

	return router
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	if len(cache) == 0 || r.URL.Query().Get("force_update") == "true" {
		updateItems()
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(cache)
	if err != nil {
		log.Printf("could not encode events into json, error: %+v", err)
	}
}

func serve() {
	// start listening
	port := getEnv("API_PORT", "80")
	srv := &http.Server{
		Handler:      createAPIConfig(),
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("API server started to listing on port :%s", port)
	log.Fatal(srv.ListenAndServe())
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
