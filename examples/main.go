package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	config "github.com/inagornyi/go-config"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir = filepath.Clean(filepath.Join(dir, ""))

	cfg, err := config.NewConfig(Config{}, dir, "config.yml")
	if err != nil {
		log.Fatal(err)
	}

	bs, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bs))
}
