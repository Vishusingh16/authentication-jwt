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
	"github.com/vishusingh16/authentication-jwt/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)


 var userCollection *mongo.Collection  = database.OpenCollection(database.Client, "user")
 var validate = validator.New()

 func HashPassword()

 func VerifyPassword()

 func Signup()gin.HandlerFunc{

	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return

		}
		validationErr :=validate.Struct(user)
		 if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		 }
		 count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		 defer cancel()
		 if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError,gin.H{"error" : "error while checking email" })
		 }
		 	count , err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
			defer cancel()
			if err != nil{
				log.Panic(err)
					c.JSON(http.StatusInternalServerError, gin.H{"error" : "error occured while checking phone number"})
			}
			if count>0{
				c.JSON(http.StatusInternalServerError, gin.H{"error" :"user already exists"} )
			}
			user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	}
 }


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
	    err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)

		defer cancel()
		if err!= nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error})
			return 
		}
		c.JSON(http.StatusOK, user)
		
	}
 }
