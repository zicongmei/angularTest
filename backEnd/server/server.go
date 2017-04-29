package server

import (
	"github.com/zicongmei/angularTest/backEnd/loadConfig"
	"net/http"
)

func Start(Configs loadConfig.ConfigStruct) {

	http.Handle("/", http.FileServer(http.Dir(Configs.Server.FrontendPath)))
	http.ListenAndServe(":"+Configs.Server.Port, nil)
}
