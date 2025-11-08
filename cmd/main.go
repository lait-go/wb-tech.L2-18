package main

import (
	"calendar/config"
	"calendar/pkg/httpserver"
	"calendar/pkg/logger"
	"fmt"
)

func main(){
	cfg, err := config.New()
	if err != nil{
		fmt.Println(err)
	}

	log := logger.New(cfg.Logger)
	log.Info("LOGGER STARTED")

	_, err = httpserver.New(cfg.Http)
	if err != nil{
		log.Error(err.Error())
	}

	log.Info("SERVER STARTED")

	fmt.Println(cfg)
}