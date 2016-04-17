package adjust

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/codegangsta/cli"

	"gopkg.in/yaml.v2"
)

const urlScheme = "https"
const urlHost = "api.adjust.com"
const configFilename = ".adjustrc"

type Settings struct {
	ConfigFilename string
	URLScheme      string
	URLHost        string
}

func NewSettings() *Settings {
	configPath := getenv("ADJUST_CLI_CONFIG_PATH", os.Getenv("HOME"))

	return &Settings{
		ConfigFilename: fmt.Sprintf("%s/%s", configPath, configFilename),
		URLScheme:      getenv("ADJUST_CLI_URL_SCHEME", urlScheme),
		URLHost:        getenv("ADJUST_CLI_URL_HOST", urlHost),
	}
}

func getenv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

type Config struct {
	UserToken string
	AppTokens []string
	AppToken  string
}

func NewConfig(c *cli.Context) *Config {
	config := &Config{AppTokens: commaSeparatedParam(c, "app-tokens")}

	if c.String("user-token") != "" {
		config.UserToken = c.String("user-token")
	}

	if c.String("app-token") != "" {
		config.AppToken = c.String("app-token")
	}

	return config
}

func ReadConfig(configFilename string) *Config {
	bytes, err := ioutil.ReadFile(configFilename)
	if os.IsNotExist(err) {
		return &Config{}
	}
	if err != nil {
		Fail(err.Error())
	}

	config := make(map[string]string)
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		Fail(err.Error())
	}

	res := &Config{
		UserToken: config["user_token"],
		AppToken:  config["app_token"],
	}

	if config["app_tokens"] != "" {
		res.AppTokens = strings.Split(config["app_tokens"], ",")
	}

	return res
}

func (config *Config) WriteConfig(configFilename string) {
	persisted := ReadConfig(configFilename)

	if config.UserToken != "" {
		persisted.UserToken = config.UserToken
	}

	if config.AppToken != "" {
		persisted.AppToken = config.AppToken
	}

	if config.AppTokens != nil {
		persisted.AppTokens = config.AppTokens
	}

	w, err := os.OpenFile(configFilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		Fail(err.Error())
	}
	buf := bufio.NewWriter(w)

	if persisted.UserToken != "" {
		_, err = fmt.Fprintf(buf, "user_token: %s\n", persisted.UserToken)
	}

	if persisted.AppToken != "" {
		_, err = fmt.Fprintf(buf, "app_token: %s\n", persisted.AppToken)
	}

	if persisted.AppTokens != nil {
		_, err = fmt.Fprintf(buf, "app_tokens: %s\n", strings.Join(persisted.AppTokens, ","))
	}

	buf.Flush()
}
