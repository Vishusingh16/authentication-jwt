package main

import(
	routes "github.com/vishusingh16/authentication-jwt/routes"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main(){

	err:= godotenv.Load(".env")

	if err!=nil {
		log.Fatal("Error Loading .env file")

	}
	port := os.Getenv("PORT")

	if port==""{
			port="8000"
			}
			
			router:= gin.New()
			router.Use(gin.Logger())

			routes.AuthRoutes(router)
			routes.UserRoutes(router)

			router.GET("/api-1", func( c *gin.Context){
				c.JSON(200, gin.H{"success":"Access granted for api-1"})
			})

			router.GET("/api-2", func(c *gin.Context){
				c.JSON(200, gin.H{"success":"Access granted for api-2"})
			})

			router.Run(":"+ port)

	}
