package main

import (
	"log"

	"github.com/josepmdc/goboilerplate/cmd"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
