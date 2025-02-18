package routers

import (
	"RetailSales/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) addRoutes(e *echo.Echo) {

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.GET("/processCSVData", services.ProcessCSVData)
	e.GET("/totalRevenue", services.GetRevenue)
	e.GET("/totalRevenueByProduct", services.GetRevenueByProduct)
	e.GET("/totalRevenueByCategory", services.GetRevenueByCategory)
	e.GET("/totalRevenueByRegion", services.GetRevenueByRegion)

}
