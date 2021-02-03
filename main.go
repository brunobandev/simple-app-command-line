package main

import (
	"fmt"
	"ip-server-name-commandline/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Start..")

	application := app.Generate()
	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
