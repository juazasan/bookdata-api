package datastore

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	guuid "github.com/google/uuid"
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

// SearchByAuthor returns books by Author
func (b *Books) SearchByAuthor(author string) *[]*BookData {
	var booksFound []*BookData
	for _, book := range *b.Store {
		if strings.Contains(strings.ToLower(book.Authors), strings.ToLower(author)) {
			log.Println("Found book" + book.Title + "written by " + author)
			booksFound = append(booksFound, book)
		}
	}
	return &booksFound
}

// SearchByTitle returns books by Title
func (b *Books) SearchByTitle(title string) *[]*BookData {
	var booksFound []*BookData
	for _, book := range *b.Store {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(title)) {
			log.Println("Found book" + book.Title)
			booksFound = append(booksFound, book)
		}
	}
	return &booksFound
}

// SearchByISBN returns book by ISBN
func (b *Books) SearchByISBN(isbn string) *BookData {
	var bookFound *BookData
	for _, book := range *b.Store {
		if strings.ToLower(book.ISBN) == strings.ToLower(isbn) {
			log.Println("Found book" + book.Title + " with ISBN:" + isbn)
			bookFound = book
			break
		}
	}
	return bookFound
}

// DeleteByISBN deletes book by ISBN
func (b *Books) DeleteByISBN(isbn string) int {
	var newStore []*BookData
	returnCode := 404
	for _, book := range *b.Store {
		if strings.ToLower(book.ISBN) != strings.ToLower(isbn) {
			newStore = append(newStore, book)
		} else {
			log.Println("Removed book with ISBN " + isbn + " from the Store")
			returnCode = 200
		}
	}
	if returnCode == 404 {
		log.Println("No book with ISBN " + isbn + " found in the Store")
	}
	b.Store = &newStore
	return returnCode
}

// AddBook adds a new Book to the Store
func (b *Books) AddBook(newBook BookData) (*BookData, int) {
	var returnCode int
	if b.SearchByISBN(newBook.ISBN) == nil {
		newBook.BookID = guuid.New().String()
		bookReader := *b.Store
		log.Println(len(*b.Store))
		log.Println(len(bookReader))
		bookReader = append(bookReader, &newBook)
		log.Println(len(bookReader))
		b.Store = &bookReader
		log.Println(len(*b.Store))
		returnCode = 200
	} else {
		returnCode = 409
	}
	return &newBook, returnCode
}
