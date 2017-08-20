package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"app/shared/workers"
)

type Server struct {
	Hostname  string `json:"Hostname"`
	UseHTTP   bool   `json:"UseHTTP"`
	UseHTTPS  bool   `json:"UseHTTPS"`
	HTTPPort  int    `json:"HTTPPort"`
	HTTPSPort int    `json:"HTTPSPort"`
	CertFile  string `json:"CertFile"`
	KeyFile   string `json:"KeyFile"`
	Pool 	  int    `json:"Pool"`
}

func Run(httpHandlers http.Handler, s Server) {
	workers.InitPool(s.Pool)
	startHTTP(httpHandlers, s)
}

func startHTTP(handlers http.Handler, s Server) {
	fmt.Println(time.Now().Format("2006-01-02 03:04r:05 PM"), "Running HTTP "+httpAddress(s))

	log.Fatal(http.ListenAndServe(httpAddress(s), handlers))
}

func httpAddress(s Server) string {
	return s.Hostname + ":" + fmt.Sprintf("%d", s.HTTPPort)
}

