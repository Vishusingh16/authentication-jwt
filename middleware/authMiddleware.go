package middleware

import(
	"fmt"
	"net/http"
	helper "github.com/vishusingh16/Authentiation-jwt/helpers"
	"github.com/gin-gonic/gin"

)
func Authenticate() gin.HandlerFunc{
	return func(c *gin.Context){
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization Required")})
			c.Abort()
			return 
		}
	claims, err:=	helper.ValidateToken(clientToken)
	if err != ""{
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Invalid token")})
		c.Abort()
		return
	}
	c.Set("email", claims.Email)
	c.Set("first_name", claims.First_name)
	c.Set("last_name", claims.Last_name)
    c.Set("uid", claims.Uid)
	c.Set("user_type", claims.User_type)
	c.Next()
	}
}