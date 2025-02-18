package utils

import (
	"RetailSales/database"
	"RetailSales/models"
	"fmt"
	"log"
)

func CreateProductTable(Product models.Product) error {
	db := database.SqlConnect()

	// Create Table Query Based on Product Struct
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS Products (
        ProductID VARCHAR(100)  PRIMARY KEY,
        ProductName VARCHAR(255) NOT NULL,
        Category VARCHAR(100),
        ProductDescription TEXT,
        UnitPrice DECIMAL(10,2) NOT NULL
    );`

	// Execute Query
	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Printf("Error in creating Product Table %v", err.Error())
		return err
	}

	fmt.Println("Products table created successfully!")
	return nil

}

func CreateOrderTable(Order models.Order) error {
	db := database.SqlConnect()

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS Orders (
        OrderID INT  PRIMARY KEY,
        CustomerID VARCHAR(100) NOT NULL,
        ProductID VARCHAR(100) NOT NULL,
        DateOfSale DATE NOT NULL,
        QuantitySold INT NOT NULL,
        Discount DECIMAL(5,2) DEFAULT 0,
        TotalPrice DECIMAL(10,2) NOT NULL,
        PaymentMethod VARCHAR(50) NOT NULL,
        ShippingCost DECIMAL(10,2) NOT NULL,
        FOREIGN KEY (CustomerID) REFERENCES Customers(CustomerID) ON DELETE CASCADE,
        FOREIGN KEY (ProductID) REFERENCES Products(ProductID) ON DELETE CASCADE
    );`

	// Execute Query
	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Printf("Error in creating Order Table %v", err.Error())
		return err
	}

	fmt.Println("Orders table created successfully!")
	return nil
}

func CreateCustomerTable(customer models.Customer) error {
	db := database.SqlConnect()

	// Create Customers Table Query
	createTableQuery := `
	   CREATE TABLE IF NOT EXISTS Customers (
		   CustomerID VARCHAR(100)  PRIMARY KEY,
		   CustomerName VARCHAR(255) NOT NULL,
		   CustomerEmail VARCHAR(255) UNIQUE NOT NULL,
		   CustomerAddress TEXT NOT NULL,
		   Region VARCHAR(100)
	   );`

	// Execute Query
	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Printf("Error in creating Customer Table %v", err.Error())
		return err
	}

	fmt.Println("Customers table created successfully!")
	return nil
}

func CreateTables() {
	err := CreateProductTable(models.Product{})
	if err != nil {
		panic(err)
	}
	err = CreateCustomerTable(models.Customer{})
	if err != nil {
		panic(err)
	}
	err = CreateOrderTable(models.Order{})
	if err != nil {
		panic(err)
	}

}
