package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"MicroServiceCore/structs"
)

var serverConfigData *structs.ServerData
var statsMiddleware *structs.ExtendedStats

// Init initialises the handler package
func Init(sd *structs.ServerData, sm *structs.ExtendedStats) {
	serverConfigData = sd
	statsMiddleware = sm
}

func writeHeader(w http.ResponseWriter, status ...int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	if status != nil {
		w.WriteHeader(status[0])
	}
}

func writeBytes(w http.ResponseWriter, b []byte) {
	writeHeader(w)
	w.Write(b)
}

func writeErrorResponse(w http.ResponseWriter, e error) {
	str := fmt.Sprintf("{ \"error\": \"%s\" }", e.Error())
	writeBytes(w, []byte(str))
}

func writeObjectResponse(o interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(o)

	if err != nil {
		writeHeader(w, http.StatusInternalServerError)
	} else {
		writeBytes(w, json)
	}
}
