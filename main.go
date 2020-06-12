package gvaplug

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/piexlmax/gva-plug-rules/rules_servies"
	"github.com/piexlmax/gvaplug/global"
	"github.com/piexlmax/gvaplug/router"
)

type Plug interface {
	InitRouter([2]*gin.RouterGroup) error
	InitModel(*gorm.DB) error
}

type GvaModel struct {
	Test string `json:"test" form:"test" gorm:"comment:'GvaModel测试字段'"`
}

type GvaPlug struct {
	SomeConfig string
}

func (g GvaPlug) InitModel(db *gorm.DB) error {
	global.GVA_DB = db
	db.AutoMigrate(GvaModel{})
	return nil
}

// Routers[0] 为纯净路由  不加载用户的中间件
// Routers[1] 为非纯净路由  加载用户的中间件

func (g GvaPlug) InitRouter(Routers [2]*gin.RouterGroup) error {
	R := Routers[0].Group("")
	router.InitGVAPlugRouter(R)
	rules_servies.RegisterApi(global.GVA_DB, global.GVA_PLUS_APIS)
	return nil
}
