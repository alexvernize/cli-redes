package commands

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/urfave/cli"
)

func BuscarIPs(c *cli.Context) {
	// Aqui você pode adicionar a lógica para buscar os IPs
	// Isso pode incluir fazer uma consulta DNS ou qualquer outra lógica necessária
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if host == "" {
		host = c.String("H")
	}
	fmt.Printf("Buscando IPs para o domínio: %s\n", host)
	if erro != nil {
		fmt.Printf("Erro ao buscar IPs para o domínio %s: %v\n", host, erro)
		return
	}
	for _, ip := range ips {
		fmt.Printf("IP encontrado: %s\n", ip.String())
	}
}

func BuscarServidoresDNS(c *cli.Context) error {
	// Aqui você pode adicionar a lógica para buscar os servidores DNS
	// Isso pode incluir fazer uma consulta DNS ou qualquer outra lógica necessária
	host := c.String("host")
	if host == "" {
		return fmt.Errorf("por favor, especifique um host com --host ou --H")
	}
	servidores, erro := net.LookupNS(host)
	fmt.Printf("Buscando servidores DNS para o domínio: %s\n", host)
	if erro != nil {
		return fmt.Errorf("erro ao buscar servidores DNS para o domínio %s: %v", host, erro)
	}
	fmt.Printf("Servidores DNS para %s:\n", host)
	for _, srv := range servidores {
		fmt.Println("→", srv.Host)
	}
	return nil
}

func PortScan(c *cli.Context) error {
	host := c.String("host")
	if host == "" {
		return cli.NewExitError("Host não especificado", 1)
	}

	timeout := time.Duration(c.Int("timeout")) * time.Second
	ports := strings.Split(c.String("ports"), ",")

	for _, portStr := range ports {
		port := strings.TrimSpace(portStr)
		address := fmt.Sprintf("%s:%s", host, port)

		conn, err := net.DialTimeout("tcp", address, timeout)
		if err == nil {
			fmt.Printf("Porta %s aberta\n", port)
			conn.Close()
		}
	}
	return nil
}

func Whois(c *cli.Context) error {
	domain := c.String("domain")
	if domain == "" {
		return cli.NewExitError("Domínio não especificado", 1)
	}

	conn, err := net.Dial("tcp", "whois.iana.org:43")
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("Erro ao conectar: %v", err), 2)
	}
	defer conn.Close()

	conn.Write([]byte(domain + "\r\n"))
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}
