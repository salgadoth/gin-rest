package main

import (
	"github.com/salgadoth/gin-rest-api/database"
	"github.com/salgadoth/gin-rest-api/routes"
)

func main() {
	database.ConectaComDB()

	routes.HandleRequests()
}
