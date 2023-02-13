package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// recibir datos -> middleware? -> sacar datos -> crear request -> hacer request a bd -> recuperar datos -> sacar datos -> retornar datos

// func getQueryData (){}
// func createRequest(){}
// func get/post/Data(){}
// func extractEmails(){}
// func returnData(){}

type emailsResource struct{}

func (resource emailsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/data", resource.List) // GET /posts - Read a list of posts.
	r.Post("/", resource.Create)   // POST /posts - Create a new post.

	r.Route("/{id}", func(r chi.Router) {
		r.Use(PostCtx)
		r.Delete("/", resource.Delete)
	})

	return r
}

func (resource emailsResource) List(w http.ResponseWriter, r *http.Request) {

	Request, _ := CreateRequest(r)

	x, _ := json.Marshal(Request)

	y := string(x)

	w.Header().Set("Content-Type", "application/json")

	req, err := http.NewRequest("POST", "http://localhost:4080/api/emails/_search", strings.NewReader(y))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(body)

	return
}

func (resource emailsResource) Create(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (resource emailsResource) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/posts/"+id, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ParseResponse(response []byte) {

	var result map[string]any

	json.Unmarshal([]byte(response), &result)

	//return getEmail(result)
}

func getEmail(result map[string]any) any {

	hits := result["hits"].(any)

	println(hits.(string))
	//hitss := hits["hits"].([]any)
	// hitsss := hitss[0].(map[string]any)
	// email := hitsss["_source"].(map[string]any)

	// em, _ := json.Marshal(email)

	return hits
}
