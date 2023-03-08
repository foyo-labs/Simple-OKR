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

//go:embed web/build/*
var content embed.FS

func main() {
	configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("Error while loading config file: ", configErr)
	}

	database, databaseError := db.ConnectDatabase(*config.C)

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
	userAPI := app.InitUserAPI(database)
	groupAPI := app.InitGroupAPI(database)
	cycleAPI := app.InitCycleAPI(database)

	api := r.Group("/api")
	objectives := api.Group("/objectives")

	users := api.Group("/users")
	groups := api.Group("/groups")
	cycles := api.Group("/cycles")

	// Objectives
	objectives.POST("", objectiveAPI.Create)

	// KeyResults
	objectives.GET("/:objectiveID/keyresults", nil)
	objectives.GET("/:objectiveID/keyresults/:keyResultId", nil)

	// Users
	users.POST("/login", userAPI.Login)
	users.POST("/registion", userAPI.Create)
	users.PUT("/groups", userAPI.UpdateGroup)

	// Units: Company & Department
	groups.POST("/", groupAPI.Create)
	groups.GET("", groupAPI.ListChilds)

	cycles.POST("", cycleAPI.Create)

	// r.StaticFS("/", http.FS(content))
	err := r.Run(config.C.Http.Port)
	if err != nil {
		panic(err)
	}
}
