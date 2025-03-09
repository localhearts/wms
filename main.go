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
	r := gin.Default()

	inboundRepo := repository.InboundRepository{DB: server.DB}
	routes.RegisterInboundRoutes(r, inboundRepo)

	log.Fatal(r.Run(":8080"))

}
