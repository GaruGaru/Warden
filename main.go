package main

import (
	"github.com/x-cray/logrus-prefixed-formatter"
	log "github.com/sirupsen/logrus"
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

}
