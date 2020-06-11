package global

import (
	"github.com/jinzhu/gorm"
	"github.com/piexlmax/gva-plug-rules/rules_model"
)

var (
	GVA_DB        *gorm.DB
	GVA_PLUS_APIS []rules_model.GvaPlugApi
)
