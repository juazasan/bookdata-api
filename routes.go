package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/juazasan/bookdata-api/datastore"
)

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	limit, err := getLimitParam(r)
	skip, err := getSkipParam(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "invalid datatype for parameter"}`))
		return
	}
	data := books.GetAllBooks(limit, skip)
	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func getLimitParam(r *http.Request) (int, error) {
	limit := 0
	queryParams := r.URL.Query()
	l := queryParams.Get("limit")
	if l != "" {
		val, err := strconv.Atoi(l)
		if err != nil {
			return limit, err
		}
		limit = val
	}
	return limit, nil
}

func getSkipParam(r *http.Request) (int, error) {
	skip := 0
	queryParams := r.URL.Query()
	l := queryParams.Get("skip")
	if l != "" {
		val, err := strconv.Atoi(l)
		if err != nil {
			return skip, err
		}
		skip = val
	}
	return skip, nil
}

func searchByAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	booksbyAuthor := books.SearchByAuthor(params["author"])
	json.NewEncoder(w).Encode(booksbyAuthor)
	return
}

func searchByTitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	booksbyTitle := books.SearchByTitle(params["title"])
	json.NewEncoder(w).Encode(booksbyTitle)
	return
}

func searchByISBN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	book := books.SearchByISBN(params["isbn"])
	json.NewEncoder(w).Encode(book)
	return
}

func deleteByISBN(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	returnCode := books.DeleteByISBN(params["isbn"])
	w.WriteHeader(returnCode)
	return
}

func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook datastore.BookData
	_ = json.NewDecoder(r.Body).Decode(&newBook)
	book, returnCode := books.AddBook(newBook)
	if returnCode == 200 {
		json.NewEncoder(w).Encode(book)
	} else {
		w.WriteHeader(returnCode)
	}
	return
}
