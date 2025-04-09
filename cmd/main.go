package main

import (
	"fmt"

	"github.com/Ulpio/xuiter.git/internal/config"
)

func main() {
	config.ConnectDB()
	config.AutoMigrate()
	fmt.Println("Banco Conecta com sucesso")
}
