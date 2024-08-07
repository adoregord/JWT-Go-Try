package main

import (
	"jwt-try/internal/provider"
	"jwt-try/internal/provider/routes"
)

func main() {
	db, err := provider.DBConnection()
	if err != nil {
		return
	}
	defer db.Close()

	routes.SetupRoutes(db).Run("127.0.0.1:8080")
}
