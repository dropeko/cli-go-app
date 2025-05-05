package main

import (
	"command-line-app/app"
	"log"
	"os"
)

func main() {
	appCLI := app.Gerar()

	if err := appCLI.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
