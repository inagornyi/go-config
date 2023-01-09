package main

type Config struct {
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"app"`
}
