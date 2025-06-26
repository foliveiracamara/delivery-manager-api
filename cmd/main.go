// @title Delivery Manager API
// @version 1.0
// @description API para gerenciamento de envios de pacotes por diferentes transportadoras
// @host localhost:5000
// @BasePath /

package main

import (
	_ "github.com/foliveiracamara/delivery-manager-api/docs"
	"github.com/foliveiracamara/delivery-manager-api/internal/cmd"
)

func main() {
	cmd.NewRoot().Execute()
}
