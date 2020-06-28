package auth

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/irvanherz/aldebaran/internal/models/user"
	"golang.org/x/crypto/bcrypt"
)

type ResponseData struct {
	Code    string      `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type SigninData struct {
	Token    string      `json:"token"`
	UserData interface{} `json:"userData,omitempty"`
}

type AuthClaims struct {
	User user.User `json:"user"`
	Test int       `json:"test"`
	jwt.StandardClaims
}

type LoginRequestData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func generate_token(u *user.User) (string, error) {
	claims := AuthClaims{
		*u,
		100,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    "Aldebaran",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("DEFILATIFAH"))
}

func Signin(c *gin.Context) {
	var data LoginRequestData
	c.BindJSON(&data)

	res, err := user.ReadByEmail(data.Username)
	if err == "" && res != nil {
		compare_err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(data.Password))
		if compare_err != nil {
			c.JSON(http.StatusInternalServerError, ResponseData{
				Code:    "AUTH_FAILED",
				Message: "Invalid password",
			})
		} else {
			token, token_err := generate_token(res)
			if token_err != nil {
				c.JSON(http.StatusInternalServerError, ResponseData{
					Code:    "SERVER_ERROR",
					Message: "Server getting an error while generating your login token",
				})
			} else {
				res.Password = ""
				c.JSON(http.StatusOK, ResponseData{
					Code: "SUCCESS",
					Data: SigninData{
						Token:    token,
						UserData: res,
					},
				})
			}
		}
	} else if err == "" && res == nil {
		c.JSON(http.StatusUnauthorized, ResponseData{
			Code:    "AUTH_FAILED",
			Message: "Invalid username",
		})
	} else {
		c.JSON(http.StatusInternalServerError, ResponseData{
			Code:    err,
			Message: "Failed reading database",
		})
	}
}

func Signup(c *gin.Context) {
	var data user.User
	c.BindJSON(&data)
	hashed_password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	data.Password = string(hashed_password)
	res, err := user.Write(data)
	if err == "" {
		res.Password = ""
		c.JSON(200, ResponseData{
			Code:    "SUCCESS",
			Message: "Signup success",
			Data:    res,
		})
	} else {
		c.JSON(200, ResponseData{
			Code:    "ERROR",
			Message: err,
		})
	}
}
