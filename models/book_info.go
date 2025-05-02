package models

type BookInfo struct {
	BookName          string `json:"book_name"`
	BookCategory      string `json:"book_category"`
	BookPublishedDate string `json:"book_published_date"`
	BookAuthor        string `json:"book_author"`
	BookDescription   string `json:"book_description"`
	BookEdition       string `json:"book_edition"`
	BookLanguage      string `json:"book_language"`
	BookPrice         string `json:"book_price"`
	BookPublisher     string `json:"book_publisher"`
}
