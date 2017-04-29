package loadConfig

import (
	"encoding/json"
	"io/ioutil"
)

type ServerStruct struct {
	Port        string `json:"port"`
	FronendPath string `json:"frontEnd"`
}

type ConfigStruct struct {
	Server ServerStruct `json:"server"`
}

var Congigs ConfigStruct

func Load(file string) error {
	bData, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bData, &Congigs); err != nil {
		return err
	}
	return nil
}
