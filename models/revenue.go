package models

// New struct to hold product details + total revenue
type ProductRevenue struct {
	ProductID    string  `json:"product_id"`
	ProductName  string  `json:"product_name"`
	UnitPrice    float64 `json:"unit_price"`
	TotalRevenue float64 `json:"total_revenue"`
}

// New struct to hold category-wise revenue

type CategoryRevenue struct {
	Category     string  `json:"category"`
	TotalRevenue float64 `json:"total_revenue"`
}

// New struct to hold region-wise revenue
type RegionRevenue struct {
	Region       string  `json:"region"`
	TotalRevenue float64 `json:"total_revenue"`
}
