package domain

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// Config is a struct for app configurations
type Config struct {
	// JwtSigningKey string `json:"JWT_SIGNING_KEY"`
	// Port          string `json:"PORT"`
	// MongoDBName   string `json:"MONGO_DB_NAME"`
	// DB   *DBConfig
	// Host *HostConfig
	// JWT  *JWTConfig
	config map[string]interface{}
	lock   sync.RWMutex
}

func convertKeysToStrings(m map[interface{}]interface{}) (map[string]interface{}, error) {
	n := make(map[string]interface{})

	// ITERATE THROUGH CONFIG KEY/VALUE PAIRS
	for k, v := range m {
		// ASSERT ALL KEYS ARE STRINGS
		str, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("config key is not a string")
		}

		if vMap, ok := v.(map[interface{}]interface{}); ok {
			var err error
			v, err = convertKeysToStrings(vMap)
			if err != nil {
				return nil, err
			}
		}

		n[str] = v
	}

	return n, nil

}

// SetFromBytes sets the internal config based on a byte array of YAML
func (c *Config) SetFromBytes(data []byte) error {
	var rawConfig interface{}
	if err := yaml.Unmarshal(data, &rawConfig); err != nil {
		return err
	}

	untypedConfig, ok := rawConfig.(map[interface{}]interface{})
	if !ok {
		return fmt.Errorf("config is not a map")
	}

	config, err := convertKeysToStrings(untypedConfig)
	if err != nil {
		return err
	}

	// LOCK MAP BEFORE WRITING TO HANDLE CONCURRENCY ISSUES
	c.lock.Lock()
	defer c.lock.Unlock()

	c.config = config
	return nil
}

// Get returns configurations for serviceName
func (c *Config) Get(serviceName string) (map[string]interface{}, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	// GET BASE CONFIG
	a, ok := c.config["base"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("base config is not a map")
	}

	// IF NO CONFIG IS DEFINED FOR THE SERVICE
	if _, ok = c.config[serviceName]; !ok {
		// RETURN THE BASE CONFIG
		return a, nil
	}

	b, ok := c.config[serviceName].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("service %q is not a map", serviceName)
	}

	// MERGE THE MAPS WITH THE SERVICE CONFIG TAKING PRESEDENCE
	config := make(map[string]interface{})
	for k, v := range a {
		config[k] = v
	}
	for k, v := range b {
		config[k] = v
	}

	return config, nil
}

func getConfigPath() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	filename := []string{"config.", env}
	return strings.Join(filename, "")
	// filename := []string{"config/", "config.", env, ".json"}
	// _, dirname, _, _ := runtime.Caller(0)
	// filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	// return filePath
}

func initViper() (Config, error) {
	viper.SetConfigName(getConfigPath()) // Set path to config file for environment
	viper.AddConfigPath(".")             // Search the root directory for the configuration file
	err := viper.ReadInConfig()          // Find and read the config file
	if err != nil {                      // Handle errors reading the config file
		return Config{}, err
	}
	viper.SetDefault("PORT", "8081") // Sets defualt value for PORT is PORT is not defined

	var config Config
	err = viper.Unmarshal(&config) // Unmarhsall the data read from config into constant struct
	// fmt.Println(config.DB.Name)
	return config, err
}

func readConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}

// New is used to generate a configuration instance which will be passed around the codebase
// func New() (*Config, error) {
// 	v1, err := readConfig(getConfigPath(), *Config{})
// 	if err != nil {
// 		panic(fmt.Errorf("error reading config: %v", err))
// 	}
// 	// config := Config{}
// 	var config Config
// 	err = viper.Unmarshal(&config)
// 	// constants, err := initViper()
// 	fmt.Println(config)
// 	return config, err
// 	// config.Constants = constants
// 	// if err != nil {
// 	// 	return &config, err
// 	// }

// 	// return &config, err
// }
