package adjust

import (
	"fmt"
	"os"
	"testing"
)

func fail(t *testing.T, headline string, actual string, expected string) {
	t.Error(headline, fmt.Sprintf("\nActual: %s\n", actual), fmt.Sprintf("\nExpect: %s\n", expected))
}

func cleanFixtures() {
	_ = os.Remove("fixtures/config_write.yaml")
}

func assertConfig(t *testing.T, config *Config, expectations map[string]string) {
	if expectations["user_token"] != "" && config.UserToken != expectations["user_token"] {
		fail(t, "Couldn't read/write config", config.UserToken, expectations["user_token"])
	}

	if expectations["app_token"] != "" && config.AppToken != expectations["app_token"] {
		fail(t, "Couldn't read/write config", config.AppToken, expectations["app_token"])
	}

	if expectations["app_token1"] != "" && config.AppTokens[0] != expectations["app_token1"] {
		fail(t, "Couldn't read/write config", config.AppTokens[0], expectations["app_token1"])
	}

	if expectations["app_token2"] != "" && config.AppTokens[1] != expectations["app_token2"] {
		fail(t, "Couldn't read/write config", config.AppTokens[1], expectations["app_token2"])
	}

	if expectations["app_token2"] != "" && len(config.AppTokens) != 2 {
		t.Error("Wrong size of app expectations")
	}
}

func TestReadConfig(t *testing.T) {
	config := ReadConfig("fixtures/config_read.yaml")

	expectations := map[string]string{
		"user_token": "user-token",
		"app_token":  "app-token",
		"app_token1": "app-token1",
		"app_token2": "app-token2",
	}

	assertConfig(t, config, expectations)
}

func TestWriteConfig(t *testing.T) {
	defer cleanFixtures()

	config := &Config{UserToken: "ivan", AppToken: "metodi", AppTokens: []string{"ivan", "metodi"}}

	config.WriteConfig("fixtures/config_write.yaml")

	written := ReadConfig("fixtures/config_write.yaml")

	expectations := map[string]string{
		"user_token": "ivan",
		"app_token":  "metodi",
		"app_token1": "ivan",
		"app_token2": "metodi",
	}

	assertConfig(t, written, expectations)
}

func TestWriteConfigOnlyAppToken(t *testing.T) {
	defer cleanFixtures()

	config := &Config{AppToken: "metodi"}

	config.WriteConfig("fixtures/config_write.yaml")

	written := ReadConfig("fixtures/config_write.yaml")

	expectations := map[string]string{
		"user_token": "",
		"app_token":  "metodi",
	}

	assertConfig(t, written, expectations)
	if written.AppTokens != nil {
		t.Error("AppTokens should be empty. Was", written.AppTokens)
	}
}

func TestWriteConfigOnlyAppTokens(t *testing.T) {
	defer cleanFixtures()

	config := &Config{AppTokens: []string{"metodi", "stefan"}}

	config.WriteConfig("fixtures/config_write.yaml")

	written := ReadConfig("fixtures/config_write.yaml")

	expectations := map[string]string{
		"user_token": "",
		"app_token":  "",
		"app_token1": "metodi",
		"app_token2": "stefan",
	}

	assertConfig(t, written, expectations)
}
