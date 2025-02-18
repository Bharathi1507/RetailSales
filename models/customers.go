package models

// Customer Struct
type Customer struct {
	CustomerID      string `json:"customer_id"`
	CustomerName    string `json:"customer_name"`
	CustomerEmail   string `json:"customer_email"`
	CustomerAddress string `json:"customer_address"`
	Region          string `json:"region"`
}
