package models

import "time"

type CustomerBookInventory struct {
	CustomerID    string    `json:"customer_id"` // Partition Key
	BookID        string    `json:"book_id"`
	OrderID       string    `json:"order_id"`
	PurchaseDate  time.Time `json:"purchase_date"` // Clustering Key
	ModeOfPayment string    `json:"mode_of_payment"`
	OrderStatus   string    `json:"order_status"`

	// Collections for books, using SET to ensure uniqueness
	ListCurrentBooks  []string `json:"list_current_books"`  // Set of Book IDs (unique, no duplicates)
	ListPreviousBooks []string `json:"list_previous_books"` // Set of Book IDs (unique, no duplicates)

	// Counters for the number of books
	CurrentBooksCount  int `json:"current_books_count"`
	PreviousBooksCount int `json:"previous_books_count"`
}
