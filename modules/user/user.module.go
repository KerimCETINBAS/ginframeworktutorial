package user

import "github.com/gin-gonic/gin"

func UserModule(r *gin.RouterGroup) {

	c := UserController()
	// stack routers
	r.GET("/", c.GetUsers)
	r.POST("/", c.CreateUser)
}
