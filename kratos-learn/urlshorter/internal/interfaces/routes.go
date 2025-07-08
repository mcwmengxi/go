package interfaces

import "github.com/gin-gonic/gin"

func RegisterHTTPServer(uc *UserUseCase) *gin.Engine {
	r := gin.Default()
	rootGrp := r.Group("/api")
	{
		rootGrp.GET("/", uc.Register)
	}
	return r
}
