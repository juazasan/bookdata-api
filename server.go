package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juazasan/bookdata-api/datastore"
)

var (
	books datastore.BookStore
)

func init() {
	books = &datastore.Books{}
	books.Initialize()
}

func main() {
	r := mux.NewRouter()
	log.Println("bookdata api")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	})
	api.HandleFunc("/books", getAllBooks).Methods(http.MethodGet)
	api.HandleFunc("/books/author/{author}", searchByAuthor).Methods(http.MethodGet)
	api.HandleFunc("/books/title/{title}", searchByTitle).Methods(http.MethodGet)
	api.HandleFunc("/book/isbn/{isbn}", searchByISBN).Methods(http.MethodGet)
	api.HandleFunc("/book/isbn/{isbn}", deleteByISBN).Methods(http.MethodDelete)
	api.HandleFunc("/book", addBook).Methods(http.MethodPost)
	log.Fatalln(http.ListenAndServe(":8080", r))
}
