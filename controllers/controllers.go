package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
)

func HashPassword(password string) string {}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeOut(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJson; err != nil {
			c.JSON{http.StatusBadRequest, gin.H{"error": err.Error()}}
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}
		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
			return
		}

		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})

		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this phone number is laready in use"})
			return
		}
	}
}

func Login() gin.HandlerFunc {}

func ProductViewerAdmin() gin.HandlerFunc {}

func SearchProduct() gin.HandlerFunc {}

func SearchProductByQuery() gin.HandlerFunc {}
