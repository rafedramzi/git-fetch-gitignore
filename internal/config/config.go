package config

import (
	"strings"
	"time"

	"fmt"
	"log"

	"github.com/go-playground/validator/v10"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	GitRepositorySource = "GIT"
	UrlSource           = "URL"
)

type Config struct {
	DefaultRepository string        `validate:"required"`
	CacheDir          string        `validate:"dir"`
	ExpireDuration    time.Duration `validate:"required"`
	Sources           []Source      `validate:"required,dive,unique=Name"`
}

type Source struct {
	Name   string `validate:"required"`
	Kind   string `validate:"required,fi_src_kind"`
	Source string `validate:"required"` // TODO: Valiate legit git url
}

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("fi_src_kind", validateSupportedSourceKind)
}

func validateSupportedSourceKind(fl validator.FieldLevel) bool {
	str := strings.ToUpper(fl.Field().String())
	if str != GitRepositorySource && str != UrlSource {
		return false
	}
	return true
}

// Validate Config data
func (c *Config) Validate() error {
	err := validate.Struct(c)

	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			return err
		}
		// TODO: Should be handled at test instead
		panic(err)
	}
	return nil
}

func Load(configFile string, config *Config) error {
	viper.SetConfigName("fetch-ignore")
	viper.AddConfigPath("$HOME/.config/fetch-ignore/")

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigType("TOML")
		viper.SetConfigFile("config.toml")
	}
	// Add Default Config Path
	// Viper Read in
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("Config file does not exists: %s", err)
		}
		return fmt.Errorf("Failed to Parse Configuration: %s", err)
	}
	if err := viper.Unmarshal(config); err != nil {
		return fmt.Errorf("Failed to unmarshal config: %s", err)
	}
	setDefaultProperty(config)
	if err := config.Validate(); err != nil {
		return fmt.Errorf("Failed to validate config: %s", err)
	}
	return nil
}

func setDefaultProperty(config *Config) {
	if config.CacheDir == "" {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalln("Unable to retrieve user home directory: ", err.Error())
		}
		config.CacheDir = home
	}
}
