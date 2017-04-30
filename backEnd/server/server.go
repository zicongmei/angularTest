package server

import (
	"encoding/json"
	"github.com/zicongmei/angularTest/backEnd/loadConfig"
	"net/http"
	"path/filepath"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("received: " + r.URL.Path))
}

type UserRequest struct {
	User     string `json:"user"`
	Password string `json:"pwd"`
}

func authenticateHandler(w http.ResponseWriter, r *http.Request) {
	var ur UserRequest
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&ur)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("tried to login with " + ur.User))
}

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	var Configs *loadConfig.ConfigStruct = &loadConfig.GlobalConfig
	http.Redirect(w, r, "https://127.0.0.1:"+Configs.Server.HttpsPort+r.RequestURI, http.StatusMovedPermanently)
}

func Start() {
	var Configs *loadConfig.ConfigStruct = &loadConfig.GlobalConfig
	frontendPAth, err := filepath.Abs(Configs.Server.FrontendPath)
	if err != nil {
		panic(err)
	}
	frontendPAth += "/"
	http.Handle("/", http.FileServer(http.Dir(frontendPAth)))
	http.HandleFunc("/request/", requestHandler)
	http.HandleFunc("/authenticate/", authenticateHandler)
	go http.ListenAndServeTLS(":"+Configs.Server.HttpsPort,
		Configs.Server.HttpsCert, Configs.Server.HttpsKey, nil)

	http.ListenAndServe(":"+Configs.Server.HttpPort,
		http.HandlerFunc(redirectToHttps))
}
