package app

import (
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de servidor na internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs publicos na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Usage: "Host desejado (Default: google.com.br)",
					Value: "google.com.br",
				},
			},
			Action: buscarIps,
		},
		{
			Name:  "server",
			Usage: "Busca o servidor em que o dominio informado está hospedado",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Usage: "Host desejado (Default: google.com.br)",
					Value: "google.com.br",
				},
			},
			Action: buscarServer,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")
	log.Println("Host pesquisado:", host)

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		log.Println(ip)
	}
}

func buscarServer(c *cli.Context) {
	host := c.String("host")
	log.Println("Host pesquisado:", host)

	servers, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		log.Println(server.Host)
	}
}
