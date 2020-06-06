package router

import (
	"github.com/gin-gonic/gin"
	"github.com/piexlmax/gvaplug/api"
)

func InitGVAPlugRouter(Router *gin.RouterGroup) {
	gvaPlug := Router.Group("gvaPlug")
	{
		gvaPlug.GET("testGvaPlug", api.GetGvaTest)
	}
}
