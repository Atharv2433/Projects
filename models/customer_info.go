package models

type CustomerInfo struct {
	CustomerName             string `json:"customer_name"`
	CustomerMembershipStatus string `json:"customer_membership_status"`
	CustomerAge              string `json:"customer_age"`
	CustomerMobile           string `json:"customer_mobile"`
	CustomerEmail            string `json:"customer_email"`
	CustomerDeliveyStatus    string `json:"delivery_status"`
	CustomerGender           string `json:"customer_gender"`
}
