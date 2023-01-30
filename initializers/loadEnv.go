package initializers

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	TokenSecret    		string        	`mapstructure:"TOKEN_SECRET"`
	TokenExpiresIn 		time.Duration 	`mapstructure:"TOKEN_EXPIRED_IN"`
	TokenMaxAge    		int           	`mapstructure:"TOKEN_MAXAGE"`
	MongoUri    		string			`mapstructure:"MONGO_URI"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}








