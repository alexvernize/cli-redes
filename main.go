package main

import (
	"cli-redes/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("App cli-redes iniciando...\n")

	aplicação := app.Gerar()
	if err := aplicação.Run(os.Args); err != nil {
		log.Fatal(err)
		// Se você quiser imprimir uma mensagem de erro personalizada, descomente a linha abaixo
		// e comente a linha acima
		//fmt.Println("Erro ao executar a aplicação:", err)
	}
	fmt.Println("\nAplicação encerrada com sucesso")
}
