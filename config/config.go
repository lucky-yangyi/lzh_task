package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type config struct {
	Type string `json:"type"`
	Host string `json:"host"`
	Root string `json:"root"`
	Pwd  string `json:"pwd"`
	Name string `json:"name"`
	Port string `json:"port"`
}

var Config *config

func init() {
	fmt.Println(os.Getwd())
	f, e := os.Open("./conf/app.json")
	defer f.Close()
	if e != nil {
		panic(e)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	con := new(config)
	if err = json.Unmarshal(b, con); err != nil {
		panic(err)
	}
	Config = con
}
