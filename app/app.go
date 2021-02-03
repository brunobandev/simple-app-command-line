package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate returns an command line application to be executed.
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "IP & Server Names"
	app.Usage = "Find IP`s and Server Names on Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Find IP`s on internet",
			Flags:  flags,
			Action: findIPs,
		},
		{
			Name:   "servers",
			Usage:  "Find server names on internet",
			Flags:  flags,
			Action: findServerNames,
		},
	}

	return app
}

func findIPs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findServerNames(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
