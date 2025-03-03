package config

import (
	"github.com/spf13/viper"
	"log"
)

type Cases struct {
	Prefix string `yaml:"prefix"`
	Pass   bool   `yaml:"pass"`
}

type Config struct {
	AppId    string           `yaml:"appId"`
	Secret   string           `yaml:"secret"`
	BizId    string           `yaml:"bizId"`
	RealCard string           `yaml:"realCard"`
	Cases    map[string]Cases `yaml:"cases"`
}

var Cfg Config

func New(path, name string) {

	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("viper read in config error", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalln("viper unmarshal error", err)
	}
}
