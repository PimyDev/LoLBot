package config

import (
	"../file_utils"
	"github.com/alexbyk/panicif"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Token string `yaml:"token"`
	ChannelID string `yaml:"channel_id"`
}

func LoadConfig() Config{
	var config Config
	filename := "config.yml"
	if file_utils.FileExists(filename) {
		file, err := ioutil.ReadFile(filename)
		panicif.Err(err)
		err = yaml.Unmarshal(file, &config)
		panicif.Err(err)
	} else {
		_, err := os.Create(filename)
		panicif.Err(err)
	}
	return config
}