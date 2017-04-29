package loadConfig

import (
	"encoding/json"
	"io/ioutil"
)

type ServerStruct struct {
	HttpPort     string `json:"httpPort"`
	FrontendPath string `json:"frontEndPath"`
}

type ConfigStruct struct {
	Server ServerStruct `json:"server"`
}

func Load(file string) (ConfigStruct, error) {
	var configs ConfigStruct
	bData, err := ioutil.ReadFile(file)
	if err != nil {
		return configs, err
	}
	if err := json.Unmarshal(bData, &configs); err != nil {
		return configs, err
	}
	return configs, err
}
