package game

import (
	_ "github.com/sigsant/nonogramas_resolver/util/logger"
)

type nonogram struct {
	grid  [][]int
	hints [][]int
}

//Por ejemplo, si se rellenara con un 5x5 el resultado
//inicial es el siguiente array multidmensional:
// &{[[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]}

// ó para visualizarlo mejor:

// [0 0 0 0 0]
// [0 0 0 0 0]
// [0 0 0 0 0]
// [0 0 0 0 0]
// [0 0 0 0 0]

func createArrMultidimensional(rows, columns int) [][]int {
	multiarr := make([][]int, rows)
	for i := range multiarr {
		multiarr[i] = make([]int, columns)
	}
	return multiarr
}

// CreateGrid crea un duplicado del casillero que se rellenará posteriormente para incluir las guías de dibujo
func CreateGrid(rows, columns int) *nonogram {
	return &nonogram{
		grid: createArrMultidimensional(rows, columns),
		//Las guias siempre seran un 5x5
		hints: createArrMultidimensional(5, 5),
	}
}
