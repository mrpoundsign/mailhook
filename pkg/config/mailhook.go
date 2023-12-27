package config

import (
	"fmt"
	"os"

	"github.com/mrpoundsign/mailhook/pkg/mail"
	"gopkg.in/yaml.v2"
)

type Auth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type MailHook struct {
	Port  int         `yaml:"port" default:"1025"`
	Host  string      `yaml:"host" default:"localhost"`
	Auth  Auth        `yaml:"auth"`
	Hooks []mail.Hook `yaml:"hooks"`
}

func NewMailHook(path string) (*MailHook, error) {
	// read yaml file

	// Read the YAML file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Parse the YAML data into a MailHook struct
	var mailhook MailHook
	err = yaml.Unmarshal(data, &mailhook)
	if err != nil {
		return nil, err
	}

	return &mailhook, nil
}

func (m *MailHook) ListenString() string {
	return fmt.Sprintf("%s:%d", m.Host, m.Port)
}
