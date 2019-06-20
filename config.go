package main

import "gopkg.in/yaml.v2"
import "io/ioutil"

var configInstance *AppConfig

type AppConfig struct {
	Path struct {
		Home  string
		Store string
	}
	Remote struct {
		URL    string `yaml:"url"`
		Branch string
		Token  string
		User   string
		Email  string
	}
	Files struct {
		Home     []string `yaml:",flow"`
		Absolute []string `yaml:",flow"`
	}
}

func config(path string) (*AppConfig, error) {
	if nil != configInstance {
		return configInstance, nil
	}

	data, err := ioutil.ReadFile(path)
	if nil != err {
		return nil, err
	}

	tmp := &AppConfig{}

	if err := yaml.Unmarshal([]byte(data), tmp); nil != err {
		return nil, err
	}

	configInstance = tmp

	return configInstance, nil
}
