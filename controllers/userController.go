package controllers

import (
	"Authentication-jwt/helpers"
	helper "Authentication-jwt/helpers"
	"Authentication-jwt/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vishusingh16/authentication-jwt/database"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)


 var userCollection *mongo.Collection  = database.OpenCollection(database.Client, "user")
 var validate = validator.New()

 func HashPassword()

 func VerifyPassword()

 func Signup()


 func Login()

 func GetUsers()

 func GetUser() gin.HandlerFunc {
	return func(c *gin.Context){
		userId := c.Params("user_id")
		
		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return

		}


		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		
	}
 }
