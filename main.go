package main

import (
	"fmt"

	"github.com/sigsant/nonogramas_resolver/game"
	// "github.com/sigsant/nonogramas_resolver/util/logger"
)

func main() {
	// logger.InfoLogger.Print("Inicializando paquete Main")
	testGrid := game.Grid{Columns: 5, Rows: 5}
	fmt.Printf(testGrid.String())

}
