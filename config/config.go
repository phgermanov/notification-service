package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

// Settings represents the configuration settings for the application.
type Settings struct {
	Port            string        `mapstructure:"PORT"`
	SlackWebhookURL string        `mapstructure:"SLACK_WEBHOOK_URL"`
	RetryDuration   time.Duration `mapstructure:"RETRY_DURATION"`
}

// LoadConfig loads the configuration settings from various sources.
func LoadConfig() Settings {
	var AppConfig Settings

	// Set the configuration file name and type
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")

	// Add the current directory as a search path for the configuration file
	viper.AddConfigPath(".")

	// Automatically read environment variables with matching names
	viper.AutomaticEnv()

	// Read the configuration from file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// Unmarshal the configuration data into the AppConfig struct
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Error unmarshaling config: %s", err)
	}

	fmt.Println("Config loaded successfully")
	return AppConfig
}
