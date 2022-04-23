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

type cashierApi struct {
	usecase usecase.CashierUseCase
}

func (api *cashierApi) GetAllCashier() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataMeta apprequest.Meta

		dataMeta.Limit, _ = strconv.Atoi(c.Query("limit"))
		dataMeta.Skip, _ = strconv.Atoi(c.Query("skip"))
		dataCashier, respMeta, err := api.usecase.GetAllCashier(dataMeta)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", gin.H{
			"cashiers": dataCashier,
			"meta":     respMeta,
		}))
	}
}

func (api *cashierApi) SearchCashierById() gin.HandlerFunc {
	return func(c *gin.Context) {
		cashierId, _ := strconv.Atoi(c.Param("cashierid"))
		resp, err := api.usecase.SearchCashierById(cashierId)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", resp))
	}
}

func (api *cashierApi) CreateCashier() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data apprequest.CashierRequest
		err := c.ShouldBindJSON(&data)
		fmt.Println(data)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		dataCreate, err := api.usecase.CreateCashier(data)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", dataCreate))
	}
}

func (api *cashierApi) UpdateCashier() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data apprequest.CashierUpdateRequest
		cashierId, _ := strconv.Atoi(c.Param("cashierid"))
		err := c.ShouldBindJSON(&data)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		errUpdate := api.usecase.UpdateCashier(cashierId, data.Name)
		if errUpdate != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(errUpdate.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", ""))
	}
}

func (api *cashierApi) DeleteCashier() gin.HandlerFunc {
	return func(c *gin.Context) {
		cashierId, _ := strconv.Atoi(c.Param("cashierid"))
		err := api.usecase.DeleteCashierUseCase(cashierId)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", ""))
	}
}

func NewCashierApi(routeGroup *gin.RouterGroup, usecase usecase.CashierUseCase) {
	api := &cashierApi{usecase}

	routeGroup.GET("", api.GetAllCashier())
	routeGroup.GET("/:cashierid", api.SearchCashierById())
	routeGroup.POST("", api.CreateCashier())
	routeGroup.PUT("/:cashierid", api.UpdateCashier())
	routeGroup.DELETE("/:cashierid", api.DeleteCashier())
}
