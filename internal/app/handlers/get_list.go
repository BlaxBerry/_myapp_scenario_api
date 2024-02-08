package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	myModels "myapp_scenario_api/internal/app/models"
	constants "myapp_scenario_api/internal/pkg/constants"
)

// 获取列表
//
//	 Paths:
//		 GET api/list
//		 GET api/list?fields=[key,key,...]
//
//	 Examples:
//		 GET api/list
//		 GET api/list?fields=id,title,updateAt,pushedAt,data
func GetList(ctx *gin.Context) {
	// ...

	_, hasFields := ctx.GetQueryArray("fields")
	if !hasFields {
		ctx.JSON(http.StatusOK, myModels.ApiResponseData{
			Code:  http.StatusOK,
			Error: nil,
			Data: []any{
				constants.MockMessageNode,
				constants.MockHTMLNode,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, myModels.ApiResponseData{
		Code:  http.StatusOK,
		Error: nil,
		Data:  []any{},
	})
}
