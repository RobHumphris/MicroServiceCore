package handlers

import (
	"log"
	"net/http"
)

// Health handler returns the Extended Statistics for the server
func Health(w http.ResponseWriter, r *http.Request) {
	if serverConfigData.Debug {
		log.Println("Handler: Health Request received.")
	}
	st := statsMiddleware.Data(serverConfigData)
	writeObjectResponse(st, w)
	if serverConfigData.Debug {
		log.Println("Handler: Health Response sent.")
	}
}
