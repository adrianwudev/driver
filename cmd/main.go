package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// handlers
	// userHandler := handlers.NewUserHandler(myDb)

	// routes
	// ticket create
	// ticket update
	// ticket select
	// router.GET("/", GreetingHandler)
	// router.POST("/api/v1/user/login", userHandler.Login)

	return router
}

func main() {
	fmt.Println("driver")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// routers
	router := setupRouter()
	err := router.Run(":8091")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server started on port 8091")
}
