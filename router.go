package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
)

const searchAddress = "http://localhost:4080/api/emails/_search"
const deleteEndpoint = "http://localhost:4080/api/emails/_doc/"

type emailsResource struct{}

func (resource emailsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", resource.List)
	r.Post("/search", resource.Search)

	return r
}

func makeRequest(w http.ResponseWriter, r *http.Request, request Request) {
	requestObject, err := json.Marshal(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	println(string(requestObject))
	req, err := http.NewRequest("POST", searchAddress, strings.NewReader(string(requestObject)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.SetBasicAuth(os.Getenv("ZINC_FIRST_ADMIN_USER"), os.Getenv("ZINC_FIRST_ADMIN_PASSWORD"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
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

func (resource emailsResource) List(w http.ResponseWriter, r *http.Request) {
	request, err := CreateRequest(r, "alldocuments")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	makeRequest(w, r, request)
}

func (resource emailsResource) Search(w http.ResponseWriter, r *http.Request) {

	request, err := CreateRequest(r, "match")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	makeRequest(w, r, request)
}
