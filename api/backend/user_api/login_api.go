package user_api

import (
	"blog/utils/response"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserLogin(c *gin.Context) {
	response.SuccessWithMessage("UserLogin", c)
}
