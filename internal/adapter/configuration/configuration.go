package configuration

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ListenPort    string `envconfig:"APP_PORT"`
	DBEndpoint    string `envconfig:"DB_ENDPOINT"`
	DBUser        string `envconfig:"DB_USER"`
	DBPassword    string `envconfig:"DB_PASSWORD"`
	DBName        string `envconfig:"DB_NAME"`
	DBPort        string `envconfig:"DB_PORT"`
	RedisEndpoint string `envconfig:"REDIS_ENDPOINT"`
	RedisPort     string `envconfig:"REDIS_PORT"`
	RedisDB       int    `envconfig:"REDIS_DB"`
	RedisPassword string `envconfig:"REDIS_PASSWORD"`
}

var globalConfig Config

func GetMySQLUri() string {
	connectTemplate := "%s:%s@tcp(%s:%s)/%s"
	return fmt.Sprintf(
		connectTemplate,
		globalConfig.DBUser,
		globalConfig.DBPassword,
		globalConfig.DBEndpoint,
		globalConfig.DBPort,
		globalConfig.DBName,
	)
}

func GetPostgresUri() string {
	// "host=myhost port=myport user=gorm dbname=gorm password=mypassword"
	connectTemplate := "host=%s port=%s user=%s dbname=%s password=%s sslmode=disable"
	return fmt.Sprintf(
		connectTemplate,
		globalConfig.DBEndpoint,
		globalConfig.DBPort,
		globalConfig.DBUser,
		globalConfig.DBName,
		globalConfig.DBPassword,
	)
}

func GetRedisAddr() string {
	connectTemplate := "%s:%s"
	return fmt.Sprintf(
		connectTemplate,
		globalConfig.RedisEndpoint,
		globalConfig.RedisPort,
	)
}

func Load() {
	envconfig.MustProcess("", &globalConfig)
}
func Get() Config {
	return globalConfig
}
