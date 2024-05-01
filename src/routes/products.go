package route

import (
	"cocean.com/src/handlers"
	"cocean.com/src/middlewares"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProductsRoutes(g *echo.Group, db *gorm.DB) {
	prodGrp := g.Group("/products")

	prodGrp.Use(middlewares.AuthMiddleware)

	prodGrp.POST("/:id" /*Id of the business*/, func(c echo.Context) error {
		return handlers.CreateProduct(c, db)
	})
	prodGrp.GET("/all", func(c echo.Context) error {
		return handlers.GetAll(c, db)
	})

	


	
} 