package initer

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DbUrI     string `mapstructure:"MONGODB_URI"`
	RedisUrI  string `mapstructure:"REDIS_URI"`
	RedisPass string `mapstructure:"REDIS_PASSWORD"`
	RedisDb   int    `mapstructure:"REDIS_DB"`

	Port string `mapstructure:"PORT"`

	JwtSecret  string `mapstructure:"JWT_SECRET"`
	JwtExpHour int    `mapstructure:"JWT_EXP_HOUR"`
	JwtIssuer  string `mapstructure:"JWT_ISSUER"`

	EmailFrom string `mapstructure:"EMAIL_FROM"`
	SmtpHost  string `mapstructure:"SMTP_HOST"`
	SmtpUser  string `mapstructure:"SMTP_USER"`
	SmtpPass  string `mapstructure:"SMTP_PASS"`
	SmtpPort  int    `mapstructure:"SMTP_PORT"`

	Origin  string `mapstructure:"CLIENT_ORIGIN"`
	BaseUrl string `mapstructure:"BASE_URL"`
}

var AppConfig *Config

func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return err
	}

	return nil
}

func InitConfig(path string) {
	if err := LoadConfig(path); err != nil {
		log.Fatal("Error loading config: ", err)
	}
}
