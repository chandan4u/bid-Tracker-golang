package lib

import (
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
)

var configDirName = "./config/"

// GetDefaultConfigDir : get deafult config directory
func GetDefaultConfigDir() (string, error) {
	var configDirLocation string
	homeDir, err := homedir.Dir()

	if err != nil {
		return "", err
	}

	switch runtime.GOOS {
	case "linux":
		// Use the XDG_CONFIG_HOME variable if it is set, otherwise
		// $HOME/.config/example
		xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
		if xdgConfigHome != "" {
			configDirLocation = xdgConfigHome
		} else {
			configDirLocation = filepath.Join(homeDir, ".config", configDirName)
		}

	default:
		// On other platforms we just use $HOME/.example
		hiddenConfigDirName := "." + configDirName
		configDirLocation = filepath.Join(homeDir, hiddenConfigDirName)
	}

	return configDirLocation, nil
}

// Config : configuration
type Config struct {
	General GeneralOptions
	Keys    map[string]map[string]string
}

// GeneralOptions : general options
type GeneralOptions struct {
	DefaultURLScheme       string
	FormatJSON             bool
	Insecure               bool
	PreserveScrollPosition bool
	Timeout                Duration
}

// Duration : duration
type Duration struct {
	time.Duration
}

// UnmarshalText : text marshal
func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

// DefaultConfig : default config
var DefaultConfig = Config{
	General: GeneralOptions{
		DefaultURLScheme:       "https",
		FormatJSON:             true,
		Insecure:               false,
		PreserveScrollPosition: true,
		Timeout: Duration{
			Duration: 1 * time.Minute,
		},
	},
}

// LoadConfig : load config file
func LoadConfig() (*Config, error) {
	stage := "development"
	if os.Getenv("stage") != "" {
		stage = os.Getenv("stage")
	}
	configFile := configDirName + stage + ".toml"
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	conf := DefaultConfig
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
