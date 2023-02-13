package main

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	SearchType  string   `json:"search_type"`
	Query       Query    `json:"query"`
	From        float64  `json:"from"`
	Max_results float64  `json:"max_results"`
	Source      []string `json:"_source"`
}

type Query struct {
	Term       string `json:"term"`
	Start_time string `json:"start_time"`
	End_time   string `json:"end_time"`
}

func CreateRequest(r *http.Request) (Request, error) {

	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return Request{}, err
	}

	println(data["searchTerm"].(string))

	if err != nil {
		return Request{}, err
	}

	query := Query{
		Term:       data["searchTerm"].(string),
		Start_time: "2022-06-02T00:00:00.000Z",
		End_time:   "2023-12-02T15:28:31.894Z",
	}

	var str []string

	request := Request{
		SearchType:  "match",
		Query:       query,
		From:        data["page"].(float64),
		Max_results: data["elementsPerPage"].(float64),
		Source:      str,
	}

	return request, nil
}
