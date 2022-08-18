package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
	JWT JWTConfig
}

type JWTConfig struct {
	SecretKey string
	Issure    string
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func init() {
	viper.SetDefault("api.port", ":4000")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.pass"),
		Name:     viper.GetString("db.name"),
	}

	cfg.JWT = JWTConfig{
		SecretKey: viper.GetString("jwt.secretkey"),
		Issure:    viper.GetString("jwt.issure"),
	}

	return nil
}

func GetDBConfig() DBConfig {
	return cfg.DB
}

func GetAPIConfig() string {
	return cfg.API.Port
}

func GetJWTConfig() JWTConfig {
	return cfg.JWT
}
