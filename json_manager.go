package main

import (
	"encoding/json"
	"net/http"
)

const dateFormat = "2006-01-02T15:04:05.000Z"

type Request struct {
	SearchType  string   `json:"search_type"`
	Query       Query    `json:"query"`
	From        float64  `json:"from"`
	Max_results float64  `json:"max_results"`
	Source      []string `json:"_source"`
}

type Query struct {
	Term string `json:"term"`
}

func CreateRequest(r *http.Request, searchType string) (Request, error) {

	var data map[string]any
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return Request{}, err
	}

	query := Query{
		Term: data["searchTerm"].(string),
	}

	//var str []string
	request := Request{
		SearchType:  searchType,
		Query:       query,
		From:        data["page"].(float64),
		Max_results: data["elementsPerPage"].(float64),
		Source:      nil,
	}

	return request, nil
}
