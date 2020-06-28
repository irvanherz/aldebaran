package auth

import (
	"github.com/gin-gonic/gin"
	authC "github.com/irvanherz/aldebaran/internal/controllers/auth"
)

func Setup(r *gin.RouterGroup) {
	s := r.Group("/signin")
	s.POST("/", authC.Signin)

	t := r.Group("/signup")
	t.POST("/", authC.Signup)
}
