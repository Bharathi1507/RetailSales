package utils

import (
	"RetailSales/database"
	"RetailSales/models"
	"fmt"
)

func InsertCustomer(customerDetails models.Customer) error {
	db := database.SqlConnect()
	query := `INSERT IGNORE INTO Customers (CustomerID, CustomerName, CustomerEmail, CustomerAddress, Region) 
              VALUES (?, ?, ?, ?, ?);`

	_, err := db.Exec(query, customerDetails.CustomerID, customerDetails.CustomerName, customerDetails.CustomerEmail, customerDetails.CustomerAddress, customerDetails.Region)
	if err != nil {
		return fmt.Errorf("error inserting customer: %v", err)
	}
	return nil
}
