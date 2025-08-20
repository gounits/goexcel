package parse_test

import (
	"fmt"
	"testing"

	"github.com/gounits/goexcel/parse"
)

func TestXLSX(t *testing.T) {
	xlsx := parse.XLSX{}
	load, err := xlsx.Load("../data/test.xlsx")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(load)
}
