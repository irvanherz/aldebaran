package auth

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/irvanherz/aldebaran/internal/data"
	"github.com/irvanherz/aldebaran/internal/models/user"
)

type AuthClaims struct {
	User user.User `json:"user"`
	Test int       `json:"test"`
	jwt.StandardClaims
}

func CheckToken(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	println(authHeader)
	authTokens := strings.Split(authHeader, " ")
	if len(authTokens) != 2 {
		println("1")
		c.AbortWithStatusJSON(http.StatusUnauthorized, data.ResponseData{
			Code:    "INVALID_TOKEN",
			Message: "Please provide valid login credential",
		})
	} else if strings.Compare(string(authTokens[0]), "Bearer") != 0 {
		println("2")
		c.AbortWithStatusJSON(http.StatusUnauthorized, data.ResponseData{
			Code:    "INVALID_TOKEN",
			Message: "Please provide valid login credential",
		})
	} else {
		parsedToken, err := jwt.ParseWithClaims(string(authTokens[1]), &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("DEFILATIFAH"), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, data.ResponseData{
				Code:    "INVALID_TOKEN",
				Message: err.Error(),
			})
		} else {
			claims, ok := parsedToken.Claims.(*AuthClaims)
			if ok && parsedToken.Valid {
				c.Set("userData", claims)
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, data.ResponseData{
					Code:    "INVALID_TOKEN",
					Message: "Oops, parsing token failed",
				})
			}
		}
	}
}

func ReadOne(c *gin.Context) {
	userData, exist := c.Get("userData")
	if !exist {
		c.AbortWithStatusJSON(http.StatusInternalServerError, data.ResponseData{
			Code:    "VALIDATION_ERROR",
			Message: "User data not exist",
		})
	} else {
		user := userData.(*AuthClaims)
		println(user.User.Name)
	}

}
