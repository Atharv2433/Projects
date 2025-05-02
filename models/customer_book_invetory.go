package models

import "time"

type CustomerBookInventory struct {
	// CustomerID    string    `json:"customer_id"` // Partition Key
	// BookID        string    `json:"book_id"`
	// OrderID       string    `json:"order_id"`
	PurchaseDate      time.Time `json:"purchase_date"` // Clustering Key
	ModeOfPayment     string    `json:"mode_of_payment"`
	BookStatus        string    `json:"order_status"`
	CustomerName      string    `json:"customer_name"`
	BookName          string    `json:"book_name"`
	ListCurrentBooks  []string  `json:"list_current_books"`
	ListPreviousBooks []string  `json:"list_previous_books"`
	TotalAmount       string    `json:"total_amount"`
	CustomerGender    string    `json:"customer_gender"`
	BookQuantity      string    `json:"book_quantity"`
}
