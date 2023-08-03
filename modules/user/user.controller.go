package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kerimcetinbas/gogintut/models"
)

// user controller objesi
type userController struct {
	userService IUserService
}

// user controller interface
type IUserController interface {
	GetUsers(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
}

// constructer
func UserController() IUserController {
	return &userController{
		userService: UserService(),
	}
}

// usercontroller.GetUsers()
func (c *userController) GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.userService.GetUsers())
}

func (c *userController) CreateUser(ctx *gin.Context) {

	user := models.User{}

	if err := ctx.BindJSON(&user); err != nil {
		panic("user could not be bind")
	}

	c.userService.CreateUser(user)
}

func (c *userController) GetUserById(ctx *gin.Context) {

	var response models.UserResponse
	param, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	id := uint(param)
	data := c.userService.GetUserById(id)

	fmt.Println("before serialize :", response.Email)
	response.FromUser(data)
	fmt.Println("after serialize :", response.Email)

	ctx.JSON(http.StatusOK, response)

}
