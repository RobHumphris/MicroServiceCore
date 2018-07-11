package structs

import (
	"errors"
	"log"
	"os"
)

func getPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return ":80", errors.New("no port defined. Use export PORT=value to set. Defaulting to localhost:80")
	}
	return port, nil
}

func getExecutablePath() (string, error) {
	path := os.Getenv("EXEPATH")
	if path == "" {
		return "./", errors.New("no executable path defined. Use export EXEPATH=value to set. Defaulting to './'")
	}
	return path, nil
}

func getMetricsServerType() (string, error) {
	mst := os.Getenv("METRICS_SERVER_TYPE")
	if mst == "kafka" || mst == "redis" {
		return mst, nil
	}
	return "", errors.New("metrics server type incorrectly set: Use METRICS_SERVER_TYPE to set either redis or kafka, no metrics will be logged")
}

func getMetricsServerEndpoint() (string, error) {
	mep := os.Getenv("METRICS_ENDPOINT")
	if mep == "" {
		return mep, errors.New("metrics endpoint not set, no metrics will be logged")
	}
	return mep, nil
}

func getDebugState() bool {
	debug := os.Getenv("DEBUG") == "1"
	if debug {
		log.Printf("Debug Messages are On")
	}
	return debug
}
