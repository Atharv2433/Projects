package models

type CustomerInfo struct {
	CustomerID                string `json:"customer_id"`                 // Unique ID for customer
	CustomerName              string `json:"customer_name"`               // Customer's name
	CustomerMembershipStatus  string `json:"customer_membership_status"`  // Membership status (Active/Inactive)
	CustomerAddr              string `json:"customer_addr"`               // Customer's address
	CustomerMobile            string `json:"customer_mobile"`             // Customer's mobile number
	CustomerEmail             string `json:"customer_email"`              // Customer's email
	CustomerGSTNumber         string `json:"customer_gst_number"`         // GST (Goods and Services Tax) number
	CustomerUsername          string `json:"customer_username"`           // Username for login
	CustomerPassword          string `json:"customer_password"`           // Password for login
	CustomerPreferredLanguage string `json:"customer_preferred_language"` // Preferred language for communication
	CustomerGender            string `json:"customer_gender"`             // Gender of the customer (Male/Female/Other)
}
