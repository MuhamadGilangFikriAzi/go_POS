package config

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"gopos.com/m/authenticator"
	"os"
	"time"
)

type Config struct {
	ConfigToken  authenticator.TokenConfig
	ConfigServer *ConfigServer
	*ConfigDatabase
}

type ConfigDatabase struct {
	mysqlConn   string
	configRedis *ConfigRedis
}

type ConfigServer struct {
	Url  string
	Port string
}

type ConfigRedis struct {
	Address  string
	Password string
	Db       int
}

func newTokenConfig() authenticator.TokenConfig {
	//duration, _ := strconv.Atoi(GetConfigValue("JWTDURATION"))
	return authenticator.TokenConfig{
		AplicationName:      "GO POS",
		JwtSignatureKey:     "P@ssw0rd",
		JwtSignatureMethod:  jwt.SigningMethodHS256,
		AccessTokenDuration: 60 * time.Minute,
	}
}

func newServerConfig() *ConfigServer {
	return &ConfigServer{
		os.Getenv("SERVER_URL"),
		os.Getenv("SERVER_PORT"),
	}
}

func (c *ConfigDatabase) MysqlConn() string {
	return c.mysqlConn
}

func (c *ConfigDatabase) RedisConfig() *ConfigRedis {
	return c.configRedis
}

func ReadConfigFile(configName string) {
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the Config file
	if err != nil {             // Handle errors reading the Config file
		panic(fmt.Errorf("Fatal error Config file: %w \n", err))
	}
}

func GetConfigValue(configName string) string {
	ReadConfigFile("Config")
	return viper.GetString(configName)
}

func newMysqlConn() string {
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbUrl := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DBNAME")
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbUrl, dbPort, dbName)
	return dsn
}

func newRedisConn() *ConfigRedis {
	return &ConfigRedis{
		Address:  fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		Db:       0,
	}
}

func NewConfig() *Config {
	return &Config{
		ConfigToken:  newTokenConfig(),
		ConfigServer: newServerConfig(),
		ConfigDatabase: &ConfigDatabase{
			newMysqlConn(),
			newRedisConn(),
		},
	}
}
