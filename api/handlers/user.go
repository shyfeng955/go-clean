package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shyfeng955/go-clean/models"
	"github.com/shyfeng955/go-clean/service"
	"net/http"
)

type UserHandel struct {
	userService service.UserService
}

func NewUserHandel(userService service.UserService) UserHandel {
	return UserHandel{
		userService: userService,
	}
}

func (u *UserHandel) GetUserInfo(c *gin.Context) {
	req := models.UserVo{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": fmt.Sprintf("params error: %v", err),
		})
		return
	}

	info, err := u.userService.GetUserInfo(c, req.Id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": fmt.Sprintf("Internal server error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    info,
	})
}
