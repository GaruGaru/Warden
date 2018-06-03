package main

import (
	"log"
	"github.com/GaruGaru/Warden/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
