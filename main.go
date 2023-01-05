package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/faq-service/config"
	"github.com/ssoql/faq-service/internal/api"
	"github.com/ssoql/faq-service/internal/infrastructure/db"
)

func main() {
	dbClient, err := initializeDB()
	if err != nil {
		panic(err)
	}

	router := createRouter()

	api.RegisterRoutes(router, dbClient)

	if err := router.Run(config.GetPort()); err != nil {
		panic(err)
	}
}

func createRouter() *gin.Engine {
	router := gin.Default()

	return router
}

func initializeDB() (*db.ClientDB, error) {
	dbConfig := db.DatabaseOptions{
		DatabaseUser:     config.GetDbUser(),
		DatabasePassword: config.GetDbPassword(),
		DatabaseSchema:   config.GetDbSchema(),
		DatabasePort:     config.GetDbPort(),
		DatabaseHost:     config.GetDbHost(),
	}

	return db.InitializeDB(dbConfig)
}
