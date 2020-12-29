package main

import (
	"github.com/edoger/zkits-logger"
)

func main() {

	log := logger.New("meulog")

	log.Info("Teste de logs")

	log.WithField("id",1).Info("adicionado um campo id no log")
}