package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv        string `mapstructure:"APP_ENV"`
	DBSOURCE      string `mapstructure:"DB_SOURCE"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	MigrationPath string `mapstructure:"MIGRATION_PATH"`
}

func LoadConfig() *Env {
	env := Env{}

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		panic(err)
	}

	if env.AppEnv == "dev" {
		fmt.Println("Running on dev environment")
	}

	return &env

}
