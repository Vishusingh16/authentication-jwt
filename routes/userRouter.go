package routes 


import(

	controller "github.com/vishusingh16/Authentication-jwt/controllers"
	           "github.com/vishusingh16/Authentication-jwt/middleware"
	            "github.com/gin-gonic/gin"
 
)


func UserRoutes( incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate() )
	incomingRoutes.Use("/users", controller.GetUsers())
	incomingRotues.Use("/users/:user_id", controller.GetUsers())
	

}