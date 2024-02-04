package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	myModels "myapp_scenario_api/internal/app/models"
)

// 获取指定 id 对应的
//
//	 Paths:
//		 GET api/:id
//
//	 Examples:
//		 GET api/123
func GetById(ctx *gin.Context) {
	// ...

	id := ctx.Param("id")
	fmt.Println(id)

	ctx.JSON(http.StatusOK, myModels.ApiResponseData{
		Code:  http.StatusOK,
		Error: nil,
		Data:  []any{},
	})
}
