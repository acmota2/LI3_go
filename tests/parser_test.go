package tests

import (
	"bufio"
	"os"
	"testing"

	parser "LI3_go/csvParser"
)

func TestParse(t *testing.T) {
	f, _ := os.Open("./parser_test.csv")
	defer f.Close()
	bufrd := bufio.NewReader(f)

	expected := [][]string{
		{"12502185", "5351700", "5351700", "2013-08-31 00:25:41", "Initial commit"},
	}
	tester := parser.ConvertData(bufrd)

	for i, s := range tester[1] {
		if s != expected[0][i] {
			t.Errorf("Expected %s but actually got %s", expected, tester)
		}
	}
}
