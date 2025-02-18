package utils

import (
	"RetailSales/database"
	"RetailSales/models"
	"fmt"
)

func InsertOrder(orderDetails models.Order) error {
	db := database.SqlConnect()

	query := `INSERT INTO Orders (OrderID,ProductID, CustomerID, DateOfSale, QuantitySold, Discount, ShippingCost, PaymentMethod,TotalPrice) 
              VALUES (?, ?, ?, ?, ?, ?, ?,?,?);`

	_, err := db.Exec(query, orderDetails.OrderID, orderDetails.ProductID, orderDetails.CustomerID, orderDetails.DateOfSale, orderDetails.QuantitySold, orderDetails.Discount, orderDetails.ShippingCost, orderDetails.PaymentMethod, orderDetails.TotalPrice)
	if err != nil {
		return fmt.Errorf("error inserting order: %v", err)
	}
	return nil
}
