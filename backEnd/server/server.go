package server

import (
	"github.com/zicongmei/angularTest/backEnd/loadConfig"
	"net/http"
	"path/filepath"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("received: " + r.URL.Path))
}

func Start(Configs loadConfig.ConfigStruct) {
	frontendPAth, err := filepath.Abs(Configs.Server.FrontendPath)
	if err != nil {
		panic(err)
	}
	frontendPAth += "/"
	http.Handle("/", http.FileServer(http.Dir(frontendPAth)))
	http.HandleFunc("/request/", requestHandler)
	http.ListenAndServe(":"+Configs.Server.HttpPort, nil)
}
