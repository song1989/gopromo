package main

import (
	log "github.com/sirupsen/logrus"
	"gopromo/app/http/routers"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	routers.Run()
}
