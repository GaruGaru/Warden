package main

import (
	"github.com/x-cray/logrus-prefixed-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/GaruGaru/Warden/agent"
	"fmt"
	"encoding/json"
)

var (
	Name    = "Warden"
	Build   = "wip"
	Version = "0"
)

func main() {
	log.SetFormatter(&prefixed.TextFormatter{})
	log.WithFields(log.Fields{
		"version": Version,
		"build":   Build,
	}).Info("Starting " + Name)

	fetcher := agent.DefaultHostInfoFetcher{}

	info, err := fetcher.Fetch()

	if err != nil{
		panic(err)
	}

	jsn, _ := json.Marshal(info)

	fmt.Println(string(jsn))

}
