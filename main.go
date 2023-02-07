package main

import (
	_ "github.com/jolienai/swagger-ui-in-go/docs"
	"github.com/jolienai/swagger-ui-in-go/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Bookstore API
// @version         1.0
// @description     A book management service API in Go using Gin framework.
// @termsOfService  https://tos.santoshk.dev

// @contact.name   Jolienai Viegas
// @contact.url    https://github.com/jolienai
// @contact.email  jolienai.pl@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	router := routes.SetupRouter()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
