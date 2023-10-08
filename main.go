package main

import {
	"os"
	"golang-restaurant-management/database"
	"golang-restaurant-management/routes"
	"golang-restaurant-management/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
}
var foodConnection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("POST")

	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)

}
