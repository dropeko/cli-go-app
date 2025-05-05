package app

import (
	"github.com/urfave/cli"
)

// Gerar vai retornar um novo aplicativo CLI configurado para gerar código.
func Gerar() *cli.App {
	app := cli.NewApp()

	app.Name = "Gerador de Código"
	app.Usage = "Gera código para o projeto"
	app.Version = "1.0.0"

	return app
}
