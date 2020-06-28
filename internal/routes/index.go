package routes

import (
	"github.com/gin-gonic/gin"
	authR "github.com/irvanherz/aldebaran/internal/routes/auth"
	userR "github.com/irvanherz/aldebaran/internal/routes/user"
)

func Setup() *gin.Engine {
	r := gin.Default()

	//Index
	idx := r.Group("/")
	idx.GET("/", func(c *gin.Context) {
		c.String(200, "Aldebaran Server v1")
	})

	//Auth
	auth := r.Group("/auth")
	authR.Setup(auth)

	//user
	user := r.Group("/users")
	userR.Setup(user)

	return r
}
