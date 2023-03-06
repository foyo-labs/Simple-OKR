package main

import (
	"embed"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	app "github.com/laidingqing/sokr/internal"
	"github.com/laidingqing/sokr/internal/config"
	"github.com/laidingqing/sokr/internal/db"
)

//go:embed ui/build/*
var content embed.FS

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("Error while loading config file: ", configErr)
	}

	database, databaseError := db.ConnectDatabase(config)

	if databaseError != nil {
		log.Fatal("Error while connecting to database: ", databaseError)
	}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	objectiveAPI := app.InitObjectiveAPI(database)

	api := r.Group("/api")
	objectives := api.Group("/objectives")

	// Objectives
	objectives.GET("", objectiveAPI.GetAll)

	err := r.Run(config.Http.Port)
	if err != nil {
		panic(err)
	}
}
