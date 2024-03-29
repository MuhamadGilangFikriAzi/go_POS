package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopos.com/m/authenticator"
	"gopos.com/m/config"
	"gopos.com/m/delivery/api"
	"gopos.com/m/delivery/middleware"
	"gopos.com/m/manager"
)

type AppServer interface {
	Run()
}

type serverConfig struct {
	gin            *gin.Engine
	Name           string
	Port           string
	InfraManager   manager.InfraManager
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
	Config         *config.Config
	Middleware     *middleware.AuthTokenMiddleware
	ConfigToken    authenticator.Token
}

func (s *serverConfig) initHeader() {
	s.gin.Use(s.Middleware.TokenAuthMiddleware())
	s.routeGroupApi()
}

func (s *serverConfig) routeGroupApi() {
	apiTesting := s.gin.Group("/testing")
	api.NewTestingApi(apiTesting)

	apiCashier := s.gin.Group("cashiers")
	api.NewLoginApi(apiCashier, s.UseCaseManager.LoginUseCase(), s.ConfigToken)
	api.NewCashierApi(apiCashier, s.UseCaseManager.CashierUseCase())

	apiCategory := s.gin.Group("categories")
	api.NewCategoryApi(apiCategory, s.UseCaseManager.CategoryUseCase())

	apiPayment := s.gin.Group("payments")
	api.NewPaymentApi(apiPayment, s.UseCaseManager.PaymentUsecase())

	apiProduct := s.gin.Group("products")
	api.NewProductApi(apiProduct, s.UseCaseManager.ProductUseCase())
}

func (s *serverConfig) Run() {
	s.initHeader()
	s.gin.Run(fmt.Sprintf("%s:%s", s.Name, s.Port))
}

func Server() AppServer {
	ginStart := gin.Default()
	config := config.NewConfig()
	infra := manager.NewInfraManager(config.ConfigDatabase)
	repo := manager.NewRepoManager(infra.MysqlConn())
	usecase := manager.NewUseCaseManager(repo)
	configToken := infra.ConfigToken(config.ConfigToken)
	middleware := middleware.NewAuthTokenMiddleware(configToken)
	return &serverConfig{
		ginStart,
		config.ConfigServer.Url,
		config.ConfigServer.Port,
		infra,
		repo,
		usecase,
		config,
		middleware,
		configToken,
	}
}
