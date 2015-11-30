package goweb

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port     int
	Tls      bool
	CertFile string
	KeyFile  string
	ServeFiles []string
}

func LoadConfig(file string) (cfg *Config) {
	content, err := ioutil.ReadFile("conf/" + file + ".json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(content, &cfg); err != nil {
		panic(err)
	}

	return
}
