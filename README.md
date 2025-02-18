# SetUp
- Install Golang version 1.22.3
- Run go mod vendor
- Run go mod tidy
- Build go build -o \<app-name\>
- Start the application by ./\<app-name\>


# Schema Details

## Products Table

This table stores the product information.

- ProductID (Primary Key)
- ProductName (Not Null)
- Category
- ProductDescription
- UnitPrice (Not Null)

### Constraints

- Primary Key: ProductID
- Not Null: ProductName, UnitPrice

## Customers Table

This table stores the customer information.

- CustomerID (Primary Key)
- CustomerName (Not Null)
- CustomerEmail (Unique, Not Null)
- CustomerAddress (Not Null)
- Region

### Constraints

- Primary Key: CustomerID
- Unique: CustomerEmail
- Not Null: CustomerName, CustomerEmail, CustomerAddress

## Orders Table

This table stores the order information.

- OrderID (Primary Key)
- CustomerID (Foreign Key - refers to CustomerID in Customers table)
- ProductID (Foreign Key - refers to ProductID in Products table)
- DateOfSale (Not Null)
- QuantitySold (Not Null)
- Discount (Default 0)
- TotalPrice (Not Null)
- PaymentMethod (Not Null)
- ShippingCost (Not Null)
- RegionID (Foreign Key - refers to RegionID in Regions table)
- CampaignID (Foreign Key - refers to CampaignID in Campaigns table)

### Constraints

- Primary Key: OrderID
- Foreign Key:
- CustomerID refers to CustomerID in Customers table (ON DELETE CASCADE)
- ProductID refers to ProductID in Products table (ON DELETE CASCADE)
- RegionID refers to RegionID in Regions table
- CampaignID refers to CampaignID in Campaigns table
- Not Null: DateOfSale, QuantitySold, TotalPrice, PaymentMethod, ShippingCost, RegionID



# Data Refresh Mechanism

Here, I have scheduled a cron every day at 12:30 AM. 
Another approach, we can create jenkins pipeline and schedule a dcron over there to hit the api's.

### API - http://localhost:3000/ProcessCSVData
- **Query Parameters**:
    - fileName: csv file to upload the data. It should be in applications route.


# Revenue and Order Management API

This API provides endpoints to retrieve various metrics related to revenue within specified date ranges. 


#### Get Total Revenue for a Date Range
-**Url**: `http://localhost:3000/totalRevenue`
- **Method**: `GET`
- **Query Parameters**:
  - `startDate`: Start date for the date range (format: YYYY-MM-DD) (2024-02-20)
  - `endDate`: End date for the date range (format: YYYY-MM-DD) (2024-05-20)
- **Response**:
  ```json
  {
    "totalRevenue": 3307.96
    }
  
  
#### Get Total Revenue by Product: (For a date range)

- **Url** : `http://localhost:3000/totalRevenueByProduct`

- **Method**: `GET`
- **Query Parameters**:
  - `startDate`: Start date for the date range (format: YYYY-MM-DD)
  - `endDate`: End date for the date range (format: YYYY-MM-DD)
- **Response**:
  ```json
    [
  {
    "product_id": "P456",
    "product_name": "iPhone 15 Pro",
    "unit_price": 1299,
    "total_revenue": 2598
  },
  {
    "product_id": "P234",
    "product_name": "Sony WH-1000XM5 Headphones",
    "unit_price": 349.99,
    "total_revenue": 349.99
  },
  {
    "product_id": "P123",
    "product_name": "UltraBoost Running Shoes",
    "unit_price": 180,
    "total_revenue": 180
  },
  {
    "product_id": "P789",
    "product_name": "Levi's 501 Jeans",
    "unit_price": 59.99,
    "total_revenue": 179.97
  }
    ]

#### Get Total Revenue by Category: (For a date range)
- **URL**: ` http://localhost:3000/totalRevenueByCategory`
- **Method**: `GET`
- **Query Parameters**:
  - `startDate`: Start date for the date range (format: YYYY-MM-DD)
  - `endDate`: End date for the date range (format: YYYY-MM-DD)
- **Response**:
    ```json
    [
  {
    "category": "Electronics",
    "total_revenue": 2947.99
  },
  {
    "category": "Shoes",
    "total_revenue": 180
  },
  {
    "category": "Clothing",
    "total_revenue": 179.97
  }
    ]

#### Get Total Revenue by Region: (For a date range)
- **URL**: `/revenue/region`
- **Method**: `GET`
- **Query Parameters**:
  - `startDate`: Start date for the date range (format: YYYY-MM-DD)
  - `endDate`: End date for the date range (format: YYYY-MM-DD)
- **Response**: 
  ```json
    [
  {
    "region": "South America",
    "total_revenue": 2778
  },
  {
    "region": "Europe",
    "total_revenue": 349.99
  },
  {
    "region": "North America",
    "total_revenue": 179.97
  }
    ]
    
