package services

import (
	"RetailSales/database"
	"RetailSales/models"
	"RetailSales/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetRevenue(c echo.Context) error {

	start_Date := c.QueryParam("start_date")
	end_Date := c.QueryParam("end_date")

	startDate, err := utils.ParseDate(start_Date)
	if err != nil {
		return err
	}
	endDate, err := utils.ParseDate(end_Date)
	if err != nil {
		return err
	}

	db := database.SqlConnect()

	query := `SELECT SUM(TotalPrice) AS total_revenue
	FROM Orders
	WHERE DateOfSale BETWEEN ? AND ?;`

	var totalRevenue float64
	queryErr := db.QueryRow(query, startDate, endDate).Scan(&totalRevenue)
	if err != nil {
		log.Printf("Error in retrieving query for totalRevenue %v", err.Error())
		return c.String(http.StatusInternalServerError, queryErr.Error())
	}
	log.Printf("Successfully retrieved totalRevenue within the given date")
	return c.JSON(http.StatusOK, map[string]float64{
		"totalRevenue": totalRevenue,
	})

}

func GetRevenueByProduct(c echo.Context) error {

	start_Date := c.QueryParam("start_date")
	end_Date := c.QueryParam("end_date")

	startDate, err := utils.ParseDate(start_Date)
	if err != nil {
		return err
	}
	endDate, err := utils.ParseDate(end_Date)
	if err != nil {
		return err
	}

	db := database.SqlConnect()

	query := `
		SELECT 
			o.ProductID, 
			p.ProductName, 
			p.UnitPrice, 
			COALESCE(SUM(o.TotalPrice), 0) AS total_revenue
		FROM 
			Orders o
		JOIN 
			Products p ON o.ProductID = p.ProductID
		WHERE 
			o.DateOfSale BETWEEN ? AND ?
		GROUP BY 
			o.ProductID, p.ProductName, p.UnitPrice
		ORDER BY 
			total_revenue DESC;
	`

	// Execute the query
	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch data"})
	}
	defer rows.Close()

	// Store results in a slice
	var revenues []models.ProductRevenue

	// Iterate through results
	for rows.Next() {
		var revenue models.ProductRevenue
		if err := rows.Scan(&revenue.ProductID, &revenue.ProductName, &revenue.UnitPrice, &revenue.TotalRevenue); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to scan data"})
		}
		revenues = append(revenues, revenue)
	}

	// Check for errors after iteration
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error processing results"})
	}

	log.Printf("Successfully retrieved product revenue within the given date range")
	return c.JSON(http.StatusOK, revenues)

}

func GetRevenueByCategory(c echo.Context) error {

	start_Date := c.QueryParam("start_date")
	end_Date := c.QueryParam("end_date")

	startDate, err := utils.ParseDate(start_Date)
	if err != nil {
		return err
	}
	endDate, err := utils.ParseDate(end_Date)
	if err != nil {
		return err
	}

	db := database.SqlConnect()

	query := `
		SELECT 
			p.Category, 
			COALESCE(SUM(o.TotalPrice), 0) AS total_revenue
		FROM 
			Orders o
		JOIN 
			Products p ON o.ProductID = p.ProductID
		WHERE 
			o.DateOfSale BETWEEN ? AND ?
		GROUP BY 
			p.Category
		ORDER BY 
			total_revenue DESC;
	`

	// Execute the query
	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch data"})
	}
	defer rows.Close()

	// Store results in a slice
	var revenues []models.CategoryRevenue

	// Iterate through results
	for rows.Next() {
		var revenue models.CategoryRevenue
		if err := rows.Scan(&revenue.Category, &revenue.TotalRevenue); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to scan data"})
		}
		revenues = append(revenues, revenue)
	}

	// Check for errors after iteration
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error processing results"})
	}

	log.Printf("Successfully retrieved totalRevenueByCategory within the given date")
	return c.JSON(http.StatusOK, revenues)
}

func GetRevenueByRegion(c echo.Context) error {
	start_Date := c.QueryParam("start_date")
	end_Date := c.QueryParam("end_date")

	startDate, err := utils.ParseDate(start_Date)
	if err != nil {
		return err
	}
	endDate, err := utils.ParseDate(end_Date)
	if err != nil {
		return err
	}

	db := database.SqlConnect()

	// SQL Query to calculate total revenue per region
	query := `
		SELECT 
			c.Region, 
			COALESCE(SUM(o.TotalPrice), 0) AS total_revenue
		FROM 
			Orders o
		JOIN 
			Customers c ON o.CustomerID = c.CustomerID
		WHERE 
			o.DateOfSale BETWEEN ? AND ?
		GROUP BY 
			c.Region
		ORDER BY 
			total_revenue DESC;
	`

	// Execute the query
	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch data"})
	}
	defer rows.Close()

	// Store results in a slice
	var revenues []models.RegionRevenue

	// Iterate through results
	for rows.Next() {
		var revenue models.RegionRevenue
		if err := rows.Scan(&revenue.Region, &revenue.TotalRevenue); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to scan data"})
		}
		revenues = append(revenues, revenue)
	}

	// Check for errors after iteration
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error processing results"})
	}

	log.Printf("Successfully retrieved totalRevenueByRegion within the given date")
	return c.JSON(http.StatusOK, revenues)

}
