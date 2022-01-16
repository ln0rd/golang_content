package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generator() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line App by Leo"
	app.Usage = "Search IP and Servers Names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "localhost",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Search IPs from web andress",
			Flags: flags,
			Action: searchIps,
		},
		{
			Name: "server",
			Usage: "Search Servers from web andress",
			Flags: flags,
			Action: searchServer,
		},
	}

	return app
}

func searchIps(c *cli.Context)  {
	var host string = c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Printf("IP [%s] HOST [%s]\n", ip, host)
	}
}

func searchServer(c *cli.Context)  {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Printf("HOST [%s] SERVER [%s]\n", host, server.Host)
	}

}