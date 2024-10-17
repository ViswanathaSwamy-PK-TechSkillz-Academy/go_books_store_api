package main

import (
	"encoding/json"
	"fmt"
	"go_books_store_api/cmd/v6/data"
	"net/http"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Welcome to Go Lang World, you've requested: %s\n", r.URL.Path)
}

func (app *application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	data := map[string]string{
		"Status": "Healthy",
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}

func (app *application) VersionInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	data := map[string]string{
		"Environment": app.config.env,
		"Version":     version,
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}

func (app *application) getBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	books := getDummyBooks()

	js, err := json.MarshalIndent(books, "", "\t")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}

func getDummyBooks() []data.Book {
	return []data.Book{
		{
			ID:        1,
			CreatedAt: time.Now(),
			Title:     "The Darkening of Tristram",
			Published: 1998,
			Pages:     300,
			Genres:    []string{"Fiction", "Thriller"},
			Rating:    4.5,
			Version:   1,
		},
		{
			ID:        2,
			CreatedAt: time.Now(),
			Title:     "The Legecy of Deckard Cain",
			Published: 2007,
			Pages:     432,
			Genres:    []string{"Fiction", "Adventure"},
			Rating:    4.9,
			Version:   1,
		},
		{
			ID:        3,          // system generated
			CreatedAt: time.Now(), // system generated
			Title:     "The Black Soulstone",
			Version:   1, // system generated
		},
	}
}
