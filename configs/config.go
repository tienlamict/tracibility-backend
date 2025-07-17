package configs

import (
	"log"
	db "tracibility/internal/database"

	"github.com/spf13/viper"
)

type Config struct {
	Database db.Config `yaml:"database"`

	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

	Ethereum struct {
		RPCUrl          string `mapstructure:"rpc_url"`
		PrivateKey      string `mapstructure:"private_key"`
		ContractAddress string `mapstructure:"contract_address"`
	} `mapstructure:"ethereum"` // Dùng mapstructure trong trường hợp tên field không trùng với tên trong file YAML
}

var config *Config

func LoadConfig() *Config {
	if config != nil {
		return config
	}

	viper.SetConfigFile("configs/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	config = &Config{}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to unmarshal config: %v", err)
	}
	return config
}
