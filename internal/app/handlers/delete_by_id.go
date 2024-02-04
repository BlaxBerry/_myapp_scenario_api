package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	myModels "myapp_scenario_api/internal/app/models"
)

// 删除指定 id 对应的
//
//	 Paths:
//		 DELETE api/:id
//
//	 Examples:
//		 DELETE api/123
func DeleteById(ctx *gin.Context) {
	// ...

	id := ctx.Param("id")
	fmt.Println(id)

	ctx.JSON(http.StatusOK, myModels.ApiResponseData{
		Code:  http.StatusOK,
		Error: nil,
		Data:  []any{},
	})
}
