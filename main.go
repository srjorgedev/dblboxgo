package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/srjorgedev/dblboxgo/internal/db"
	handlerData "github.com/srjorgedev/dblboxgo/internal/handler/data"
	handlerUnit "github.com/srjorgedev/dblboxgo/internal/handler/unit"
	repositoryData "github.com/srjorgedev/dblboxgo/internal/repository/data"
	repositoryUnit "github.com/srjorgedev/dblboxgo/internal/repository/unit"
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

	unitRepo := repositoryUnit.NewSQLUnitRepository(cn)
	unitHandler := handlerUnit.NewUnitHandler(unitRepo)

	dataRepo := repositoryData.NewSQLDataRepository(cn)
	dataHandler := handlerData.NewDataHandler(dataRepo)

	r := gin.Default()

	unitRoutes := r.Group("/api/v1/unit")
	{
		unitRoutes.GET("/:id", unitHandler.GetUnitByID)
		unitRoutes.GET("/sum/:id", unitHandler.GetUnitSummaryByID)
		unitRoutes.GET("/sum", unitHandler.GetAllUnitSummaries)
	}

	dataRoutes := r.Group("/api/v1/data")
	{
		dataRoutes.GET("/tags", dataHandler.GetTags)
		dataRoutes.GET("/chapters", dataHandler.GetChapters)
		dataRoutes.GET("/rarities", dataHandler.GetRarities)
		dataRoutes.GET("/types", dataHandler.GetTypes)
		dataRoutes.GET("/affinities", dataHandler.GetAffinities)
	}

	go startHealthCheck()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

func startHealthCheck() {
	ticker := time.NewTicker(50 * time.Minute)
	defer ticker.Stop()

	health()

	for range ticker.C {
		health()
	}
}

func health() {
	fmt.Println("[ API ] The system is ok. ", time.Now())
}
