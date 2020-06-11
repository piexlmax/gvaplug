package router

import (
	"github.com/gin-gonic/gin"
	"github.com/piexlmax/gva-plug-rules/rules_model"
	"github.com/piexlmax/gvaplug/api"
	"github.com/piexlmax/gvaplug/global"
)

func InitGVAPlugRouter(Router *gin.RouterGroup) {
	gvaPlug := Router.Group("gvaPlug")
	{
		gvaPlug.GET("testGvaPlug", api.GetGvaTest)
	}
	global.GVA_PLUS_APIS = append(global.GVA_PLUS_APIS, rules_model.GvaPlugApi{
		Path:        "/gvaPlug/testGvaPlug",
		Method:      "GET",
		Description: "测试插件路由注册",
		ApiGroup:    "gvaPlug",
	})
}
