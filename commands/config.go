package main

import (
	"io/ioutil"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

// Config holds the PayPal ClientID and ClientSecret used to access the PayPal API
type Config struct {
        LiveURL      string `yaml:"live_url"`
        SandboxURL   string `yaml:"sandbox_url"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

// Init unmarshalls Config from YAML configuration in filename
func Init(filename string) (*Config, error) {
	defer glog.Flush()
	var config = new(Config)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	glog.V(5).Infof("read config %v\n", config)
	return config, err
}
