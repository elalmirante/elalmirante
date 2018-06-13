package config

import (
	"io/ioutil"

	"github.com/elalmirante/elalmirante/models"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Servers  map[string]models.Server `yaml:"servers"`
	Provider string                   `yaml:"provider"`
}

func GetServersFromConfigFile(path string) ([]models.Server, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config = Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	servers := make([]models.Server, 0)

	// add name as tag
	for k, v := range config.Servers {
		v.Tags = append(v.Tags, k)
		v.Name = k

		if v.Provider == "" {
			v.Provider = config.Provider
		}

		servers = append(servers, v)
	}

	err = validateConfiguration(servers)
	return servers, err
}
