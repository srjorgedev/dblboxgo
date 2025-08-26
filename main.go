package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/srjorgedev/dblboxgo/internal/db"
	handler "github.com/srjorgedev/dblboxgo/internal/handler/unit"
	repository "github.com/srjorgedev/dblboxgo/internal/repository/unit"
)

func main() {
	godotenv.Load()

	fmt.Println("Waiting for the connection...")

	cn, err := db.OpenConnection()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to the database successfully.")

	db.CreateTables(cn)

	unitRepo := repository.NewSQLUnitRepository(cn)
	unitHandler := handler.NewUnitHandler(unitRepo)

	r := gin.Default()

	unitRoutes := r.Group("/api/v1/unit")
	{
		unitRoutes.GET("/sum/:id", unitHandler.GetUnitSummaryByID)
		unitRoutes.GET("/sum", unitHandler.GetAllUnitSummaries)
	}

	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
