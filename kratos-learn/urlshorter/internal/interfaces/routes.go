package interfaces

import "github.com/gin-gonic/gin"

func RegisterHTTPServer(uc *UserUseCase) *gin.Engine {
	r := gin.Default()
	r.Any("/register", uc.Register)
	return r
}
