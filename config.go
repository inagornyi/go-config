package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func NewConfig[T any](name string) (*T, error) {
	bs, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var t T

	if err := yaml.Unmarshal(bs, &t); err != nil {
		return nil, err
	}

	return &t, nil
}
