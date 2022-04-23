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

type paymentApi struct {
	usecase usecase.PaymentUseCase
}

func (api *paymentApi) GetAllPayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataMeta apprequest.Meta
		dataMeta.Limit, _ = strconv.Atoi(c.Query("limit"))
		dataMeta.Skip, _ = strconv.Atoi(c.Query("skip"))
		data, respMeta, err := api.usecase.GetAllPayment(dataMeta)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", gin.H{
			"payments": data,
			"meta":     respMeta,
		}))
	}
}

func (api *paymentApi) SearchPaymentById() gin.HandlerFunc {
	return func(c *gin.Context) {
		detailId, _ := strconv.Atoi(c.Param("paymentId"))
		resp, err := api.usecase.DetailPayment(detailId)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", resp))
	}
}

func (api *paymentApi) CreatePayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data apprequest.PaymentRequest
		err := c.ShouldBindJSON(&data)
		fmt.Println(data)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		dataCreate, err := api.usecase.CreatePayment(data)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", dataCreate))
	}
}

func (api *paymentApi) UpdatePayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data apprequest.PaymentRequest
		updateId, _ := strconv.Atoi(c.Param("paymentId"))
		err := c.ShouldBindJSON(&data)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		errUpdate := api.usecase.UpdatePayment(updateId, data)
		if errUpdate != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(errUpdate.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", ""))
	}
}

func (api *paymentApi) DeletePayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		deleteId, _ := strconv.Atoi(c.Param("paymentId"))
		err := api.usecase.DeletePayment(deleteId)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", ""))
	}
}

func NewPaymentApi(routeGroup *gin.RouterGroup, usecase usecase.PaymentUseCase) {
	api := &paymentApi{usecase}

	routeGroup.GET("", api.GetAllPayment())
	routeGroup.GET("/:paymentId", api.SearchPaymentById())
	routeGroup.POST("", api.CreatePayment())
	routeGroup.PUT("/:paymentId", api.UpdatePayment())
	routeGroup.DELETE("/:paymentId", api.DeletePayment())
}
