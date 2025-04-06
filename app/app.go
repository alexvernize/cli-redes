package app

import (
	"cli-redes/commands"
	"fmt"

	"github.com/urfave/cli"
)

// Gerar cria e retorna uma nova instância de cli.App com comandos pré-definidos.
// Esta função é responsável por configurar a aplicação CLI, incluindo o nome, versão e comandos disponíveis.
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação CLI"
	app.Usage = "Uma aplicação com alguns recursos de rede que usa o pacote cli"
	app.Description = "Aplicação de exemplo para demonstrar o uso do pacote cli em Go com alguns recursos de rede."
	app.Author = "Alex Prado"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output, o", // -o ou --output
			Value: "text",
			Usage: "Formato de saída (text|json|yaml)",
		},
		cli.BoolFlag{
			Name:  "color, c", // -c ou --color
			Usage: "Ativar saída colorida",
		},
	}

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
			Usage: "Domínio a ser buscado",
		},
		cli.StringFlag{
			Name:  "H",
			Value: "google.com",
			Usage: "Domínio a ser buscado",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de domínio na internet, utilizar apenas um IP.",
			Flags:  flags,
			Action: commands.BuscarIPs,
		},
		{
			Name:    "dns",
			Aliases: []string{"d"},
			Usage:   "Busca o nome dos servidores na internet, utilizar apenas um dominio.",
			Flags:   flags,
			Action:  commands.BuscarServidoresDNS,
		},
		{
			Name:    "ping",
			Aliases: []string{"p"},
			Usage:   "Ping a host",
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return cli.NewExitError("Please provide a host to ping", 1)
				}
				host := c.Args().Get(0)
				fmt.Printf("Pinging %s...\n", host)
				// Aqui você pode adicionar a lógica para fazer o ping
				return nil
			},
		},
		{
			Name:    "portscan",
			Aliases: []string{"ps"},
			Usage:   "Verifica portas abertas em um host",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Usage: "Host para scanear",
				},
				cli.StringFlag{
					Name:  "ports",
					Value: "80,443,22,8080",
					Usage: "Portas para verificar (separadas por vírgula)",
				},
				cli.IntFlag{
					Name:  "timeout",
					Value: 2,
					Usage: "Timeout em segundos para cada tentativa",
				},
			},
			Action: commands.PortScan,
		},
		{
			Name:  "whois",
			Usage: "Consulta informações WHOIS de um domínio",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "domain",
					Usage: "Domínio para consulta WHOIS",
				},
			},
			Action: commands.Whois,
		},
	}

	return app
}
