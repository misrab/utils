package utils

import (
	"testing"

	"log"
)

func TestIndices(t *testing.T) {
	log.Println("Testing indices...")

	rows, cols := 12, 16

	index := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if CoordToIndex(i, j, rows, cols) != index {
				t.Errorf("CoordToIndex failed\n")
			}

			test_i, test_j := IndexToCoord(index, rows, cols)
			if test_i != i || test_j != j {
				t.Errorf("IndexToCoord failed\n")
			}

			index++
		}
	}
}