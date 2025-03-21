package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/localhearts/wms/database"
	"github.com/localhearts/wms/repository"
	"github.com/localhearts/wms/routes"
	"github.com/localhearts/wms/seeds"
)

func main() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	server := &database.Server{}
	server.Initialize(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	seeds.Load(server.DB)
	seeds.SeedUom(server.DB, "./data/uom.json")
	seeds.SeedCategory(server.DB, "./data/category.json")

	r := gin.Default()

	inboundRepo := repository.InboundRepository{DB: server.DB}
	categoryRepo := repository.NewCatRepository(server.DB)
	uomRepo := repository.NewUomRepository(server.DB)
	routes.CategoryRoutes(r, categoryRepo)
	routes.UomRoutes(r, uomRepo)
	// buatkan saya main route untuk inbounds

	routes.RegisterInboundRoutes(r, inboundRepo)

	log.Fatal(r.Run(":8080"))

}
