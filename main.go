package main

import (
	"jakeri-backend/middleware"
	"jakeri-backend/routers"
	"jakeri-backend/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadDotEnv()
}

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	routers.BuildRoutes(router)
	port, exists := os.LookupEnv("HTTP_PORT")
	if exists {
		router.Run(":" + port)
	} else {
		router.RunTLS(":"+os.Getenv("HTTPS_PORT"), os.Getenv("SSL_CERTIFICATE_PATH"), os.Getenv("SSL_KEY_PATH"))
	}
}
