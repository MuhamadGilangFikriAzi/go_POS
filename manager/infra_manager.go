package manager

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopos.com/m/authenticator"
	"gopos.com/m/config"
	"os"
)

type InfraManager interface {
	MysqlConn() *sqlx.DB
	RedisConn() (context.Context, *redis.Client)
	ConfigToken(tokenConfig authenticator.TokenConfig) authenticator.Token
}

type infraManager struct {
	mysqlConn *sqlx.DB
	redisConn *redis.Client
	ctx       context.Context
}

func (i *infraManager) MysqlConn() *sqlx.DB {
	return i.mysqlConn
}

func (i *infraManager) RedisConn() (context.Context, *redis.Client) {
	return i.ctx, i.redisConn
}

func (i *infraManager) ConfigToken(tokenConfig authenticator.TokenConfig) authenticator.Token {
	return authenticator.NewToken(tokenConfig, i.ctx, i.redisConn)
}

func NewInfraManager(configDatabase *config.ConfigDatabase) InfraManager {
	urlMysql := configDatabase.MysqlConn()
	redisConfig := configDatabase.RedisConfig()
	conn, err := sqlx.Connect("mysql", urlMysql)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})
	ctx := context.Background()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return &infraManager{
		mysqlConn: conn,
		redisConn: rdb,
		ctx:       ctx,
	}
}
