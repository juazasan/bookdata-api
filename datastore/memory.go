package datastore

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// Books is the memory-backed datastore used by the API
// It contains a single field 'Store', which is (a pointer to) a slice of loader.BookData struct pointers
type Books struct {
	Store *[]*BookData `json:"store"`
}

// Initialize is the method used to populate the in-memory datastore.
// At the beginning, this simply returns a pointer to the struct literal.
// You need to change this to load data from the CSV file
func (b *Books) Initialize() {
	booksDatabasePath := os.Getenv("BOOKS_DATABASE_FILE")
	log.Println("Loading books from: " + booksDatabasePath)
	booksDatabaseFile, err := os.Open(booksDatabasePath)
	if err != nil {
		log.Fatalln("Error reading: " + booksDatabasePath)
	}
	defer booksDatabaseFile.Close()
	r := csv.NewReader(booksDatabaseFile)

	var bookReader []*BookData

	for {
		book, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error reading: " + booksDatabasePath)
		}

		AverageRating, _ := strconv.ParseFloat(book[3], 64)
		NumPages, _ := strconv.ParseInt(book[7], 10, 32)
		Ratings, _ := strconv.ParseInt(book[8], 10, 32)
		Reviews, _ := strconv.ParseInt(book[9], 10, 32)

		bookReader = append(bookReader, &BookData{
			BookID:        book[0],
			Title:         book[1],
			Authors:       book[2],
			AverageRating: AverageRating,
			ISBN:          book[4],
			ISBN13:        book[5],
			LanguageCode:  book[6],
			NumPages:      int(NumPages),
			Ratings:       int(Ratings),
			Reviews:       int(Reviews)})
	}

	b.Store = &bookReader
}

// GetAllBooks returns the entire dataset, subjet to the rudimentary limit & skip parameters
func (b *Books) GetAllBooks(limit, skip int) *[]*BookData {
	if limit == 0 || limit > len(*b.Store) {
		limit = len(*b.Store)
	}
	ret := (*b.Store)[skip:limit]
	return &ret
}
