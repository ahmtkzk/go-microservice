package server

import (
	"github.com/labstack/echo/v4"
	"go-microservice/internal/database"
	"net/http"
)

type ServerConnection struct {
	echo echo.Echo
	pr   *database.ProductRepository
}

func NewServerConnection(pr *database.ProductRepository) *ServerConnection {
	return &ServerConnection{
		echo: *echo.New(),
		pr:   pr,
	}
}

func (s *ServerConnection) RegisterRoutes() {
	s.echo.GET("/products", s.GetAllProducts)
	s.echo.GET("/product", s.GetProductByName)
}

func (s *ServerConnection) StartServer() {
	s.echo.Start(":8080")
}

func (s *ServerConnection) GetAllProducts(ctx echo.Context) error {
	c := s.pr.GetAllProducts(ctx.Request().Context())
	return ctx.JSON(http.StatusOK, c)
}

func (s *ServerConnection) GetProductByName(ctx echo.Context) error {
	name := ctx.QueryParam("name")
	c := s.pr.GetProductByName(ctx.Request().Context(), name)
	return ctx.JSON(http.StatusOK, c)
}
