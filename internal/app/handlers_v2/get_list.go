package handlers_v2

import (
	"net/http"

	"github.com/gin-gonic/gin"

	myModels "myapp_scenario_api/internal/app/models"
)

// 获取列表
//
//	 Paths:
//		 GET api/v2/list
//		 GET api/v2/list?fields=[key,key,...]
//
//	 Paths:
//		 GET api/v2/list
//		 GET api/v2/list?fields=id,title,updateAt,pushedAt,data
func GetList(ctx *gin.Context) {
	// ...

	ctx.JSON(http.StatusOK, myModels.ApiResponseData{
		Code:  http.StatusOK,
		Error: nil,
		Data:  []any{},
	})
}
