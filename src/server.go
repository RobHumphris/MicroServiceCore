package main

import (
	"log"
	"net/http"

	"./handlers"
	"./structs"
	"github.com/gorilla/mux"
)

func setupRouting() *mux.Router {
	mux := mux.NewRouter()

	//mux.HandleFunc("/config", handleConfig).Methods("GET", "POST")
	//mux.HandleFunc("/control", handleControl).Methods("GET", "POST")
	//mux.HandleFunc("/stream-config", handleStreamConfig).Methods("POST")
	mux.HandleFunc("/health", handlers.Health).Methods("GET")
	return mux
}

func main() {
	serverConfigData := structs.NewServerData()
	extendedStats := structs.NewExtendedStats()
	handlers.Init(serverConfigData, extendedStats)
	mux := setupRouting()
	log.Printf("Transcode Server listening on port %v\n", serverConfigData.Port)
	log.Fatal(http.ListenAndServe(serverConfigData.Port, extendedStats.Handler(mux)))
}
