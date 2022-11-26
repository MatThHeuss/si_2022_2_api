package configs

import (
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver                     string `mapstructure:"DB_DRIVER"`
	DBHost                       string `mapstructure:"DB_HOST"`
	DBPort                       string `mapstructure:"DB_PORT"`
	MysqlUser                    string `mapstructure:"MYSQL_USER"`
	MysqlPassword                string `mapstructure:"MYSQL_PASSWORD"`
	MysqlDatabase                string `mapstructure:"MYSQL_DATABASE"`
	MysqlRootPassword            string `mapstructure:"MYSQL_ROOT_PASSWORD"`
	GoogleApplicationCredentials string `mapstructure:"GOOGLE_APPLICATION_CREDENTIALS"`
	WebServerPort                string `mapstructure:"WEB_SERVER_PORT"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
