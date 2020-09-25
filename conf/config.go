package conf

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config struct defines the configuration for the entire application
type Config struct {
	Port           int64          `yaml:"port"`
	Domain         string         `yaml:"domain"`
	LogConfig      LogConfig      `yaml:"logconfig"`
	Token          string         `yaml:"token"`
	Storage        string         `yaml:"storage"`
	PostgresConfig PostgresConfig `yaml:"db"`
}

// LogConfig defines the configuration for the Logger
type LogConfig struct {
	File  string `yaml:"file"`
	Debug bool   `yaml:"debug"`
}

// LoadConfig takes a command as an argument to get the command flags
// in case the user specified special settings. Then it loads the config from specified file
func LoadConfig(cmd *cobra.Command) (*Config, error) {

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return nil, err
	}

	if configFile, _ := cmd.Flags().GetString("config"); configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("$GOPATH/src/github.com/josepmdc/goboilerplate")
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := new(Config)
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
