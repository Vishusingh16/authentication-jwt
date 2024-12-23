package routes 


import(

	controller "github.com/vishusingh16/authentication-jwt/controllers"
	           "github.com/vishusingh16/authentication-jwt/middleware"
	            "github.com/gin-gonic/gin"
 
)


func UserRoutes( incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUser())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	

}