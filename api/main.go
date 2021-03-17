package main

import (
	"log"
	"net/http"

	"github.com/Paulo-Eduardo/inventoryservice/database"
	"github.com/Paulo-Eduardo/inventoryservice/product"
	"github.com/Paulo-Eduardo/inventoryservice/receipt"
	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	product.SetupRoutes(apiBasePath)
	receipt.SetupRoutes(apiBasePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
