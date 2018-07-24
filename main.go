package main

import (
	"log"

	"github.com/arschles/hd/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
