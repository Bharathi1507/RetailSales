package models

type Product struct {
	ProductID          string  `json:"product_id"`
	ProductName        string  `json:"product_name"`
	Category           string  `json:"category"`
	ProductDescription string  `json:"product_description"`
	UnitPrice          float64 `json:"unit_price"`
}
