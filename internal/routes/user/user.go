package user

import (
	"github.com/gin-gonic/gin"
	userC "github.com/irvanherz/aldebaran/internal/controllers/user"
	authM "github.com/irvanherz/aldebaran/internal/middlewares/auth"
)

func Setup(r *gin.RouterGroup) {
	s := r.Group("/")
	s.Use(authM.CheckToken)
	s.GET("/", userC.ReadMany)

	t := r.Group("/:userId")
	t.Use(authM.CheckToken)
	t.Use(authM.ReadOne)
	t.GET("/", userC.ReadOne)
}
