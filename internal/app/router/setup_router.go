package router

import (
	"github.com/gin-gonic/gin"

	handlers "myapp_scenario_api/internal/app/handlers"
	handlers_v2 "myapp_scenario_api/internal/app/handlers_v2"
)

// 设置路由组
//
//	 Paths:
//		 api/**/*
//		 api/v2/**/*
func SetupRouter(r *gin.Engine) {
	rgp := r.Group("/api")
	{
		rgp.GET("/list", handlers.GetList)
		rgp.GET("/:id", handlers.GetById)
		rgp.DELETE("/:id", handlers.DeleteById)
		rgp.PATCH("/:id", handlers.UpdateById)
	}

	v2 := rgp.Group("/v2")
	{
		v2.GET("/list", handlers_v2.GetList)
		v2.GET("/:id", handlers_v2.GetList)
		v2.DELETE("/:id", handlers_v2.GetList)
		v2.PATCH("/:id", handlers_v2.GetList)
	}

}
