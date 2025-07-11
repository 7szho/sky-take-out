package admin

import (
	"github.com/gin-gonic/gin"
	"take-out/internal/api/controller"
	"take-out/middle"
)

type CommonRouter struct{}

func (cr *CommonRouter) InitApiRouter(parent *gin.RouterGroup) {
	privateRouter := parent.Group("common")
	privateRouter.Use(middle.VerifyJWTAdmin())
	commonCtrl := new(controller.CommonController)
	{
		privateRouter.POST("upload", commonCtrl.Upload)
	}
}
