package main

import (
	"log"
	"os"
	"runtime"

	"app/shared/jsonconfig"
	"app/shared/database"
	"app/shared/server"
	"app/route"
	"app/shared/session"
	"encoding/json"
)

func init(){
	log.SetFlags(log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main(){
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)

	session.Configure(config.Session)

	database.Connect(config.Database)

	server.Run(route.LoadHTTP(), route.LoadHTTPS(), config.Server)
}

var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database  database.Info   `json:"Database"`
	//Email     email.SMTPInfo  `json:"Email"`
	//Recaptcha recaptcha.Info  `json:"Recaptcha"`
	Server    server.Server   `json:"Server"`
	Session   session.Session `json:"Session"`
	//Template  view.Template   `json:"Template"`
	//View      view.View       `json:"View"`
}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
