package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	AirqKey string `json:"airq_key"`
	FBKey   string `json:"fb_key"`
}

var (
	config Config
)

func init() {
	loadConfig("config.json")
}

func loadConfig(fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
}
