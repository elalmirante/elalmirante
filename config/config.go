package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Servers map[string]Server `yaml:servers`
}

type Server struct {
	Name string
	Host string   `yaml:"host"`
	Key  string   `yaml:"key"`
	Tags []string `yaml:"tags"`
}

func GetServersFromConfigFile(path string) ([]Server, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config = Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	servers := make([]Server, 0)

	// add name as tag
	for k, v := range config.Servers {
		v.Tags = append(v.Tags, k)
		v.Name = k

		servers = append(servers, v)
	}

	return servers, nil
}
