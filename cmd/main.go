package main

import (
	"fmt"

	"github.com/Ulpio/xuiter.git/internal/config"
	"github.com/Ulpio/xuiter.git/internal/routes"
)

func main() {
	config.ConnectDB()
	config.AutoMigrate()
	fmt.Println("Banco Conecta com sucesso")
	routes.RoutesHandler()
	fmt.Println("Servidor rodando na porta 8080")
}
