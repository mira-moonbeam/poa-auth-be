package config

import "github.com/spf13/viper"

type Config struct {
	DBHost            string `mapstructure:"DB_HOST"`
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	DBPort            string `mapstructure:"DB_PORT"`
	ApiSecret         string `mapstructure:"API_SECRET"`
	TokenHourLifespan string `mapstructure:"TOKEN_HOUR_LIFESPAN"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
