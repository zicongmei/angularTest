package loadConfig

import (
	"encoding/json"
	"io/ioutil"
)

type ServerStruct struct {
	HttpPort     string `json:"httpPort"`
	HttpsPort    string `json:"httpsPort"`
	HttpsCert    string `json:"httpsCert"`
	HttpsKey     string `json:"httpsKey"`
	FrontendPath string `json:"frontEndPath"`
}

type ConfigStruct struct {
	Server ServerStruct `json:"server"`
}

var GlobalConfig ConfigStruct

func Load(file string) error {
	bData, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bData, &GlobalConfig); err != nil {
		return err
	}
	return err
}
