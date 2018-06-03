package main

import (
	"github.com/GaruGaru/Warden/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
