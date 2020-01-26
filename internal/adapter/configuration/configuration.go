package configuration

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ListenPort string `envconfig:"APP_PORT"`
	DBEndpoint string `envconfig:"DB_ENDPOINT"`
	DBUser     string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASSWORD"`
	DBName     string `envconfig:"DB_NAME"`
	DBPort    string `envconfig:"DB_PORT"`

}

var globalConfig Config

func GetDBUri() string {
	connectTemplate := "%s:%s@tcp(%s:%s)/%s"
	return fmt.Sprintf(
		connectTemplate,
		Get().DBUser,
		Get().DBPassword,
		Get().DBEndpoint,
		Get().DBPort,
		Get().DBName,
		)
}

func Load() {
	envconfig.MustProcess("", &globalConfig)
}
func Get() Config {
	return globalConfig
}
