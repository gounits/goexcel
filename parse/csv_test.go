package parse_test

import (
	"fmt"
	"testing"

	"github.com/gounits/goexcel/parse"
)

func TestCSV(t *testing.T) {
	csv := parse.CSV{Comment: ','}
	load, err := csv.Load("../../data/test.csv")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(load)
}
