package game

import (
	"fmt"

	_ "github.com/sigsant/nonogramas_resolver/util/logger"
)

type Grid struct {
	Columns int
	Rows    int
}

func (g *Grid) String() string {
	// logger.DebugLogger.Printf("Grid creado con dimensiones %dx%d", g.Columns, g.Rows)
	return fmt.Sprintf("Grid creado con dimensiones %dx%d", g.Columns, g.Rows)
}
