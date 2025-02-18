package utils

import (
	"RetailSales/database"
	"RetailSales/models"
	"fmt"
)

func InsertProduct(productDetails models.Product) error {
	db := database.SqlConnect()

	query := `INSERT IGNORE INTO Products (ProductID, ProductName, Category, UnitPrice) 
              VALUES (?, ?, ?, ?);`

	_, err := db.Exec(query, productDetails.ProductID, productDetails.ProductName, productDetails.Category, productDetails.UnitPrice)
	if err != nil {
		return fmt.Errorf("error inserting product: %v", err)
	}
	return nil
}
