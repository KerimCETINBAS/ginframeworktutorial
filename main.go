package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kerimcetinbas/gogintut/modules/user"
)

func main() {

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello wolrd")
	})
	// stack modules
	user.UserModule(r.Group("/users"))

	r.Run(":8080")

}
