package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type emailsResource struct{}

func (resource emailsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/data", resource.List)
	r.Delete("/delete", resource.Delete)

	return r
}

func (resource emailsResource) List(w http.ResponseWriter, r *http.Request) {

	Request, err := CreateRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	requestObject, err := json.Marshal(Request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	req, err := http.NewRequest("POST", searchAddress, strings.NewReader(string(requestObject)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	req.SetBasicAuth("admin", "Complexpass#123")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (resource emailsResource) Delete(w http.ResponseWriter, r *http.Request) {
	//ToDo
}

func ParseResponse(response []byte) {

	var result map[string]any

	json.Unmarshal([]byte(response), &result)

	//return getEmail(result)
}

func getEmail(result map[string]any) any {

	hits := result["hits"].(any)
	//hitss := hits["hits"].([]any)
	// hitsss := hitss[0].(map[string]any)
	// email := hitsss["_source"].(map[string]any)

	// em, _ := json.Marshal(email)

	return hits
}
