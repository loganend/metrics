package main

import (
	"log"
	"os"
	"runtime"
	"encoding/json"

	"app/shared/jsonconfig"
	"app/shared/database"
	"app/shared/server"
	"app/route"
)

func init(){
	log.SetFlags(log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main(){
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)

	database.Connect(config.Database)

	server.Run(route.LoadHTTP(), config.Server)
}

var config = &configuration{}

type configuration struct {
	Database  database.Info   `json:"Database"`
	Server    server.Server   `json:"Server"`
}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
