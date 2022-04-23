package api

import (
	"github.com/gin-gonic/gin"
	"gopos.com/m/delivery/apprequest"
	"gopos.com/m/delivery/common_resp"
	"gopos.com/m/usecase"
	"net/http"
	"strconv"
)

type productApi struct {
	usecase usecase.ProductUseCase
}

func (api *productApi) GetAllProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataMeta apprequest.Meta
		dataMeta.Limit, _ = strconv.Atoi(c.Query("limit"))
		dataMeta.Skip, _ = strconv.Atoi(c.Query("skip"))
		data, respMeta, err := api.usecase.GetAllProduct(dataMeta)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", gin.H{
			"products": data,
			"meta":     respMeta,
		}))
	}
}

func (api *productApi) SearchProductById() gin.HandlerFunc {
	return func(c *gin.Context) {
		detailId, _ := strconv.Atoi(c.Param("productId"))
		resp, err := api.usecase.DetailProduct(detailId)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", resp))
	}
}

func (api *productApi) CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data apprequest.ProductRequest
		err := c.ShouldBindJSON(&data)
		//fmt.Println(data)
		//fmt.Println(time.Now().Format(data.Diskon.ExpiredAt))
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", data))

		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		//dataCreate, err := api.usecase.CreateProduct(data)
		//if err != nil {
		//	common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
		//	return
		//}
		//common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", dataCreate))
	}
}

func (api *productApi) UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data apprequest.ProductRequest
		updateId, _ := strconv.Atoi(c.Param("productId"))
		err := c.ShouldBindJSON(&data)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		errUpdate := api.usecase.UpdateProduct(updateId, data)
		if errUpdate != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(errUpdate.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", ""))
	}
}

func (api *productApi) DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		deleteId, _ := strconv.Atoi(c.Param("productId"))
		err := api.usecase.DeleteProduct(deleteId)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", ""))
	}
}

func NewProductApi(routeGroup *gin.RouterGroup, usecase usecase.ProductUseCase) {
	api := &productApi{usecase}

	routeGroup.GET("", api.GetAllProduct())
	routeGroup.GET("/:productId", api.SearchProductById())
	routeGroup.POST("", api.CreateProduct())
	routeGroup.PUT("/:productId", api.UpdateProduct())
	routeGroup.DELETE("/:productId", api.DeleteProduct())
}
