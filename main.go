package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rounin-rp/email-phisher/database"
	"github.com/rounin-rp/email-phisher/docs"
	"github.com/rounin-rp/email-phisher/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @version         1.0
// @description     This is a sample server for an email API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9000
// @BasePath  /

func main() {

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Failed to read .env file")
	}
	r := gin.Default()

	db := database.Connect()
	docs.SwaggerInfo.Title = "Email Phisher API"
	docs.SwaggerInfo.Version = "1.0"
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "It Works!",
		})
	})

	handlers.RegisterRoutes(r, db)
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":9000")
	if err != nil {
		log.Fatal("Failed to run at port 9000 due to ", err)
	}
}
