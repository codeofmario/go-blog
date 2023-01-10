package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog.com/goblog/internal/goblog/config"
)

func main() {
	settings := config.InitSettings()
	db := config.InitDB(settings)
	store := config.InitStore(settings)
	redis := config.InitRedis(settings)
	config.InitMigrations(db)
	config.SeedDemoData(db)

	router := gin.Default()
	InitRoutes(router, settings, db, store, redis)
	err := router.Run()

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
