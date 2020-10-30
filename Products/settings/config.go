package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/hashicorp/go-hclog"
)

type Configurations struct {
	SeviceConfig struct {
		ServerName  string `yaml:"name"`
		Url         string `yaml:"url"`
		Port        string `yaml:"port"`
		ApiBasePath string `yaml:"api-base-path"`
	} `yaml:"server"`
	GeneralConfig struct {
		LogLevel string       `yaml:"logLevel"`
		log      hclog.Logger `yaml:"-"`
	}
}

func NewConfig(l hclog.Logger) *Configurations {
	c := &Configurations{}
	c.GeneralConfig.log = l
	return c
}

func (c *Configurations) Load(filename string) error {
	c.GeneralConfig.log.Info("Loading config", "config", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		c.GeneralConfig.log.Error("Can't load configurations", "error", err)
		return err
	}

	//No nested struct supported by YAML
	err = yaml.Unmarshal(data, c)
	if err != nil {
		c.GeneralConfig.log.Error("Can't read  general configurations", "error", err)
		return err
	}
	return nil
}
