package models

import (
	"fmt"
	"strings"
)

type Server struct {
	Name     string
	Host     string   `yaml:"host"`
	Provider string   `yaml:"provider"`
	Key      string   `yaml:"key"`
	Tags     []string `yaml:"tags"`
}

func (s Server) String() string {
	tags := strings.Join(s.Tags, "\n\t")
	return fmt.Sprintf("Name: %s\nHost: %s\nProvider: %s\nKey: %s\nTags:\n\t%s",
		s.Name,
		s.Host,
		s.Provider,
		s.Key,
		tags)
}
