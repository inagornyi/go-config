package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func NewConfig[T any](t T, path, filename string) (*T, error) {
	bs, err := os.ReadFile(fmt.Sprintf("%s/%s", path, filename))
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(bs, &t); err != nil {
		return nil, err
	}

	return &t, nil
}
