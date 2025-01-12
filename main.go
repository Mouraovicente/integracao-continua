package main

import (
	"github.com/integracao-continua/database"
	"github.com/integracao-continua/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
