package binoku

import (
	"fmt"
	"strconv"
)

// BoardToKey takes a board and generates a key
func BoardToKey(board [][]int) string {
	key := ""
	for _, row := range board {
		for _, cell := range row {
			v := strconv.Itoa(cell)
			if cell < 0 {
				v = "x"
			}

			key = fmt.Sprintf("%s%s", key, v)
		}
	}

	return key
}
