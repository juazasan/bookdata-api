package datastore

// BookStore is the interface that the http methods use to call the backend datastore
// Using an interface means we could replace the datastore with something else,
// as long as that something else provides these method signatures...
type BookStore interface {
	Initialize()
	GetAllBooks(limit, skip int) *[]*BookData
	SearchByAuthor(author string) *[]*BookData
	SearchByTitle(title string) *[]*BookData
	SearchByISBN(isbn string) *BookData
	DeleteByISBN(isbn string) int
	AddBook(newBook BookData) (*BookData, int)
}

// BookData is the record structure of the books datastore
type BookData struct {
	BookID        string  `json:"book_id"`
	Title         string  `json:"title"`
	Authors       string  `json:"authors"`
	AverageRating float64 `json:"average_rating"`
	ISBN          string  `json:"isbn"`
	ISBN13        string  `json:"isbn_13"`
	LanguageCode  string  `json:"language_code"`
	NumPages      int     `json:"num_pages"`
	Ratings       int     `json:"ratings"`
	Reviews       int     `json:"reviews"`
}
