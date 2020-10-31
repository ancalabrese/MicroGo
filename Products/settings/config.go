package config

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/hashicorp/go-hclog"
)

type Configurations struct {
	SeviceConfig struct {
		ServerName         string        `yaml:"name"`
		Url                string        `yaml:"url"`
		Port               string        `yaml:"port"`
		ApiBasePath        string        `yaml:"api-base-path"`
		CORSAllowedOrigins string        `yaml:"cors-allowed-origins"`
		ApiRoutes          []string      `yaml:"api-routes"`
		IdleTimeout        time.Duration `yaml:"idle-timeout"`
		ReadTimeout        time.Duration `yaml:"read-timeout"`
		WriteTimeout       time.Duration `yaml:"write-timeout"`
	} `yaml:"server"`
	GeneralConfig struct {
		LogLevel string       `yaml:"logLevel"`
		log      hclog.Logger `yaml:"-"`
	}
}

func NewConfig(l hclog.Logger) *Configurations {
	c := &Configurations{}
	c.GeneralConfig.log = l
	c.SeviceConfig.ServerName = "Product-Api"
	c.SeviceConfig.Url = "localhost"
	c.SeviceConfig.Port = "9090"
	c.SeviceConfig.ApiBasePath = "/"
	c.SeviceConfig.IdleTimeout = 120 * time.Second
	c.SeviceConfig.WriteTimeout = 1 * time.Second
	c.SeviceConfig.ReadTimeout = 1 * time.Second
	return c
}

func (c *Configurations) Load(filename string) error {
	c.GeneralConfig.log.Info("Loading config", "config", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		c.GeneralConfig.log.Error("Can't load configurations", "error", err, "setting up defaults", "")
		return err
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		c.GeneralConfig.log.Error("Can't read  general configurations", "error", err, "setting up defaults")
		return err
	}
	return nil
}
