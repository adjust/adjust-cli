package adjust

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type Session struct {
	UserToken      string
	ConfigFilename string
}

func NewSession(configFilename string, userToken string) (*Session, error) {
	session := &Session{
		UserToken:      userToken,
		ConfigFilename: configFilename,
	}

	return session, session.persistSession()
}

func ReadSession(configFilename string) (*Session, error) {
	bytes, err := ioutil.ReadFile(configFilename)
	if err != nil {
		return nil, err
	}

	config := make(map[string]string)
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &Session{
		UserToken:      config["user_token"],
		ConfigFilename: configFilename,
	}, nil
}

func DestroySession(configFilename string) error {
	return os.Remove(configFilename)
}

func (session *Session) persistSession() error {
	config := []byte(fmt.Sprintf("user_token: %s\n", session.UserToken))

	return ioutil.WriteFile(session.ConfigFilename, config, 0644)
}

func (session *Session) isLoggedIn() bool {
	return session.UserToken != ""
}

func DefaultHeaders(userToken string) *http.Header {
	res := http.Header{}
	res.Add("Authorization", fmt.Sprintf("Token token=%s", userToken))
	res.Add("X-Adjust-CLI", "1")

	return &res
}
