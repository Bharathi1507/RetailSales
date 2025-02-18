package routers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Host    string
	Port    string
	Envvars []string
}

func (s *Server) Start() {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{

		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
	}))

	s.addRoutes(e)

	// ensure Envvars
	for _, envV := range s.Envvars {
		if val := os.Getenv(envV); val == "" {
			fmt.Printf("Environment Variable '%s' probably not set\n", envV)
		}
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", s.Host, s.Port)))
}

func NewServer() *Server {

	return &Server{
		Host:    "0.0.0.0",
		Port:    "3000",
		Envvars: []string{},
	}
}
