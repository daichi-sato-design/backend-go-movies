package main

import (
	"encoding/json"
	"net/http"
)

type AppStatus struct{
	Status		string `json:"status"`
	Environment	string `json:"environment"`
	Version		string `json:"version"`
}

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request){
	currentStatus := AppStatus{
		Status: "Available",
		Environment: app.config.env,
		Version: version,
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}