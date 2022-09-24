package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Create() *CliApp {
	c := &CliApp{cli.NewApp()}
	c.Name = "CLI url info"
	c.Usage = "Aplicação de linha de comando que devolver informações sobre uma URL."
	c.Commands = []cli.Command{
		c.cmmdGetIps(),
		c.cmmdGetServers(),
	}
	return c
}

type CliApp struct {
	*cli.App
}

func (c CliApp) cmmdGetIps() cli.Command {
	return cli.Command{
		Name:  "ip",
		Usage: "Busca os endereços IPs de URLs na intenet.",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "host"},
		},
		Action: getIps,
	}
}

func (c CliApp) cmmdGetServers() cli.Command {
	return cli.Command{
		Name:  "servers",
		Usage: "Busca os servidores onde a URLs está hospedada.",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "host"},
		},
		Action: getServers,
	}
}

func getIps(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println("ip:", ip)
	}
}

func getServers(c *cli.Context) {
	host := c.String("host")
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, sv := range servers {
		fmt.Println("server:", sv.Host)
	}
}
