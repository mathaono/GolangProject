package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//ToGenerate retorna a aplicação de linha de comando que será executada
func ToGenerate() *cli.App {
	app := cli.NewApp()

	app.Name = "App de linha de comando"
	app.Usage = "IPs e Nomes de servidores na net"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",           //Flag que será passada como parâmetro do comando "ip" com um valor atribuido; ex: --host amazon.com.br
			Value: "devbook.com.br", //Valor padrão passado para o comando caso não especifique a flag
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip", //Comando que será chamado na linha de comando
			Usage:  "Busca IPs de endereços",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "Buscando servidores",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

//Para pegar o valor padrão de IP não devemos passar a flag host, apenas chamar o comando ip
