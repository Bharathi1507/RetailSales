package services

import (
	"RetailSales/models"
	"RetailSales/utils"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
)

func ProcessCSVData(c echo.Context) error {
	err := ProcessCsvData(c.QueryParam("filePath"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Data uploaded Successfully",
	})
}

func ProcessCsvData(filepath string) error {
	fmt.Println("Uploading CSV Data")

	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("Error in opening the file %v", err)
		return err
	}
	defer file.Close()
	// Read CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("Error in reading the contents in file %v", err)
		return err
	}
	var products []models.Product
	var customers []models.Customer
	var orders []models.Order
	// Skip header row
	for i, row := range records {
		if i == 0 {
			continue
		}

		orderID, err := strconv.Atoi(row[0])
		if err != nil {
			log.Printf("Skipping row %d: invalid OrderID (%s)\n", i+1, row[0])
			continue
		}

		dateOfSale, err := time.Parse("2006-01-02", row[6])
		if err != nil {
			log.Printf("Skipping row %d: invalid DateOfSale (%s)\n", i+1, row[6])
			continue
		}

		quantitySold, err := strconv.Atoi(row[7])
		if err != nil {
			log.Printf("Skipping row %d: invalid QuantitySold (%s)\n", i+1, row[7])
			continue
		}

		unitPrice, err := strconv.ParseFloat(row[8], 64)
		if err != nil {
			log.Printf("Skipping row %d: invalid UnitPrice (%s)\n", i+1, row[8])
			continue
		}

		discount, _ := strconv.ParseFloat(row[9], 64)
		shippingCost, _ := strconv.ParseFloat(row[10], 64)

		// Calculate total price correctly
		totalPrice := float64(quantitySold) * unitPrice

		// Create struct instances
		product := models.Product{
			ProductID:   row[1],
			ProductName: row[3],
			Category:    row[4],
			UnitPrice:   unitPrice,
		}

		customer := models.Customer{
			CustomerID:      row[2],
			CustomerName:    row[12],
			CustomerEmail:   row[13],
			CustomerAddress: row[14],
			Region:          row[5],
		}

		order := models.Order{
			OrderID:       orderID,
			ProductID:     row[1],
			CustomerID:    row[2],
			DateOfSale:    dateOfSale,
			QuantitySold:  quantitySold,
			Discount:      discount,
			ShippingCost:  shippingCost,
			PaymentMethod: row[11],
			TotalPrice:    totalPrice,
		}
		products = append(products, product)
		customers = append(customers, customer)
		orders = append(orders, order)
	}
	for _, customer := range customers {
		if err := utils.InsertCustomer(customer); err != nil {
			log.Printf("Failed to insert customer (ID: %s): %v\n", customer.CustomerID, err)
		}
	}

	for _, product := range products {
		if err := utils.InsertProduct(product); err != nil {
			log.Printf("Failed to insert product (ID: %s): %v\n", product.ProductID, err)
		}
	}

	for _, order := range orders {
		if err := utils.InsertOrder(order); err != nil {
			log.Printf("Failed to insert order (ID: %d): %v\n", order.OrderID, err)
		}
	}

	log.Printf("Inserted OrderID successfully!\n")
	return nil

}

func RunCronJob() {

	c := cron.New(cron.WithSeconds()) // Use cron with seconds

	// Schedule the job at 12:30 AM (HH:MM:SS -> 00:30:00) everyday
	_, err := c.AddFunc("0 30 0 * * *", func() {
		ProcessCsvData("data.csv")
	})
	if err != nil {
		fmt.Println("Error scheduling cron job:", err)
		return
	}

	c.Start() // Start cron scheduler

	// Keep the main function running
	select {}
}
