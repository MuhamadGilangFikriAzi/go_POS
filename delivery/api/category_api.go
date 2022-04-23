package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopos.com/m/delivery/apprequest"
	"gopos.com/m/delivery/common_resp"
	"gopos.com/m/usecase"
	"net/http"
	"strconv"
)

type categoryApi struct {
	usecase usecase.CategoryUseCase
}

func (api *categoryApi) GetAllCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataMeta apprequest.Meta
		dataMeta.Limit, _ = strconv.Atoi(c.Query("limit"))
		dataMeta.Skip, _ = strconv.Atoi(c.Query("skip"))
		dataCategory, respMeta, err := api.usecase.GetAllCategory(dataMeta)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", gin.H{
			"categories": dataCategory,
			"meta":       respMeta,
		}))
	}
}

func (api *categoryApi) SearchCategoryById() gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryId, _ := strconv.Atoi(c.Param("categoryId"))
		resp, err := api.usecase.DetailCateogry(categoryId)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", resp))
	}
}

func (api *categoryApi) CreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data apprequest.CategoryRequest
		err := c.ShouldBindJSON(&data)
		fmt.Println(data)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		dataCreate, err := api.usecase.CreateCategory(data)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", dataCreate))
	}
}

func (api *categoryApi) UpdateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data apprequest.CategoryRequest
		categoryId, _ := strconv.Atoi(c.Param("categoryId"))
		err := c.ShouldBindJSON(&data)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		errUpdate := api.usecase.UpdateCategory(categoryId, data.Name)
		if errUpdate != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(errUpdate.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", ""))
	}
}

func (api *categoryApi) DeleteCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		cashierId, _ := strconv.Atoi(c.Param("categoryId"))
		err := api.usecase.DeleteCategory(cashierId)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", ""))
	}
}

func NewCategoryApi(routeGroup *gin.RouterGroup, usecase usecase.CategoryUseCase) {
	api := &categoryApi{usecase}

	routeGroup.GET("", api.GetAllCategory())
	routeGroup.GET("/:categoryId", api.SearchCategoryById())
	routeGroup.POST("", api.CreateCategory())
	routeGroup.PUT("/:categoryId", api.UpdateCategory())
	routeGroup.DELETE("/:categoryId", api.DeleteCategory())
}
