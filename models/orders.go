package models

import "time"

type Order struct {
	OrderID       int       `json:"order_id"`
	CustomerID    string    `json:"customer_id"`
	ProductID     string    `json:"product_id"`
	DateOfSale    time.Time `json:"date_of_sale"`
	QuantitySold  int       `json:"quantity_sold"`
	Discount      float64   `json:"discount"`
	TotalPrice    float64   `json:"total_price"`
	PaymentMethod string    `json:"payment_method"`
	ShippingCost  float64   `json:"shipping_cost"`
}
