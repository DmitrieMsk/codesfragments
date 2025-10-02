package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	router *gin.Engine
	*UserRepo
}

func NewUserHadnler(router *gin.Engine, userRepo *UserRepo) {
	handler := &UserHandler{
		router:   router,
		UserRepo: userRepo,
	}
	router.GET("/ping", handler.GetAll)

}

func (u *UserHandler) GetAll(c *gin.Context) {
	users := u.UserRepo.GetAll()
	fmt.Println(*users)

	c.JSON(http.StatusOK, users)
	//должен просто принимать get запрос на /users
	//to do...

}
