package models

import (
	"fmt"
	"strings"
)

type Server struct {
	Name     string
	Provider string   `yaml:"provider"`
	Key      string   `yaml:"key"`
	Tags     []string `yaml:"tags"`
}

func (s Server) String() string {
	tags := strings.Join(s.Tags, "\n\t")
	return fmt.Sprintf("\nName: %s\nProvider: %s\nKey: %s\nTags:\n\t%s\n",
		s.Name,
		s.Provider,
		s.Key,
		tags)
}
