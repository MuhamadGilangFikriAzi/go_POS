package api

import (
	"github.com/gin-gonic/gin"
	"gopos.com/m/authenticator"
	"gopos.com/m/delivery/apprequest"
	"gopos.com/m/delivery/common_resp"
	"gopos.com/m/usecase"
	"net/http"
	"strconv"
)

type loginApi struct {
	usecase     usecase.LoginUseCase
	configToken authenticator.Token
}

func (api *loginApi) GetCashierPasscode() gin.HandlerFunc {
	return func(c *gin.Context) {
		cashierId, _ := strconv.Atoi(c.Param("cashierid"))
		data, err := api.usecase.GetCashierPasscode(cashierId)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", gin.H{"passcode": data.Passcode}))
	}
}

func (api *loginApi) LoginCashier() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataLogin apprequest.LoginRequest
		cashierId, _ := strconv.Atoi(c.Param("cashierid"))
		if errBind := c.ShouldBindJSON(&dataLogin); errBind != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(errBind.Error()))
			return
		}
		data, isAvailable, err := api.usecase.LoginCashier(cashierId, dataLogin.Passcode)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		if isAvailable == 0 {
			common_resp.NewCommonResp(c).FailedResp(http.StatusUnauthorized, common_resp.FailedMessage("not register"))
			return
		}
		tokenString, errToken := api.configToken.CreateToken(data)
		if errToken != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage("Token Failed"))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("login admin", gin.H{
			"token": tokenString,
		}))
	}
}

func (api *loginApi) LogoutCashier() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataLogin apprequest.LoginRequest
		cashierId, _ := strconv.Atoi(c.Param("cashierid"))
		if errBind := c.ShouldBindJSON(&dataLogin); errBind != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(errBind.Error()))
			return
		}
		isAvailable, err := api.usecase.LogoutCashier(cashierId, dataLogin.Passcode)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		if isAvailable == 0 {
			common_resp.NewCommonResp(c).FailedResp(http.StatusUnauthorized, common_resp.FailedMessage("not register"))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", ""))
	}
}

func NewLoginApi(routeGroup *gin.RouterGroup, usecase usecase.LoginUseCase, configToken authenticator.Token) {
	api := &loginApi{
		usecase,
		configToken,
	}
	routeGroup.GET("/:cashierid/passcode", api.GetCashierPasscode())
	routeGroup.POST("/:cashierid/login", api.LoginCashier())
	routeGroup.POST("/:cashierid/logout", api.LogoutCashier())
}
