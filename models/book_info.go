package models

import (
	"time"

	"github.com/gocql/gocql"
)

type BookInfo struct {
	BookID            gocql.UUID `json:"book_id"`
	BookName          string     `json:"book_name"`
	Genre             string     `json:"genre"`
	BookAuthor        string     `json:"book_author"`
	BookPrice         string     `json:"book_price"`
	BookLanguage      string     `json:"book_language"`
	BookPublishedDate time.Time  `json:"book_published_date"`
	BookEdition       string     `json:"book_edition"`
	BookPublisher     string     `json:"book_publisher"`
	BookDescription   string     `json:"book_description"`
}
