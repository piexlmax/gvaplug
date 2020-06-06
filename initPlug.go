package gvaplug

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/piexlmax/gvaplug/global"
	"github.com/piexlmax/gvaplug/router"
)

type Plug interface {
	InitRouter(*gin.Engine) error
	InitModel(*gorm.DB) error
}

type GvaModel struct {
	Test string `json:"test" form:"test" gorm:"comment:'GvaModel测试字段'"`
}

type GvaPlug struct {
}

func (g GvaPlug) InitRouter(Router *gin.Engine) error {
	R := Router.Group("gvaPlug")
	router.InitGVAPlugRouter(R)
	return nil
}

func (g GvaPlug) InitModel(db *gorm.DB) error {
	global.GVA_DB = db
	db.AutoMigrate(GvaModel{})
	return nil
}
