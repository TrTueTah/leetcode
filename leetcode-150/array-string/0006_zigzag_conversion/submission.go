package zigzagconversion

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 || numRows > len(s) {
		return s
	}
	rows := make([]string, numRows)
	curRow := 0
	isDown := false

	for _, ch := range s {
		rows[curRow] += string(ch)
		if curRow == 0 || curRow == numRows-1 {
			isDown = !isDown
		}
		if isDown {
			curRow++
		} else {
			curRow--
		}
	}

	return strings.Join(rows, "")
}
