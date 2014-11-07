package utils

import ()

/*
	!! Indices start from 0, not 1
	Be sure to input in this format, 
	and alter output if your require.
*/

// returns 1-d index for 2-d coordinate
func CoordToIndex(i, j, rows, cols int) int {
	return i*cols + j
}

// returns 2-d coordinate for 1-d index
func IndexToCoord(index, rows, cols int) (int, int) {
	r := index/cols
	c := index%cols

	return r, c
}