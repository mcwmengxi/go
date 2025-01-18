package interfaces

import (
	"github.com/mcwmengxi/go/kratos-learn/urlshorter/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)


type UserUseCase struct {
	userService *service.UserService
	log         *log.Helper
}

func NewUserUseCase(sh *service.UserService, logger log.Logger) *UserUseCase{
	return &UserUseCase{
		userService: sh,
		log: log.NewHelper(logger),
	}
}

func (uc *UserUseCase)Register (c *gin.Context){
	c.JSON(200,gin.H{
		"msg":"OK",
	})
}
