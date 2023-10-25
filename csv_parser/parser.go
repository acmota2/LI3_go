package csv_parser

import (
	"bufio"
	"strconv"
	"strings"
)

func ConvertData(f *bufio.Reader) [][]string {
	var result [][]string = nil
	for {
		l, err := f.ReadString('\n')
		l = strings.TrimSuffix(l, "\n")
		current_line := string(l)
		if err != nil {
			break
		}
		vec := strings.Split(current_line, ";")
		result = append(result, vec)
	}
	return result
}

func ParseVec(vec string) []uint64 {
	to_transform := strings.Split(vec, "[")
	to_transform = strings.Split(to_transform[1], "]")
	to_transform = strings.Split(to_transform[0], ", ")
	var result []uint64 = nil
	for _, i := range to_transform {
		u, _ := strconv.ParseUint(i, 10, 64)
		result = append(result, u)
	}
	return result
}
