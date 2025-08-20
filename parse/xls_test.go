package parse_test

import (
	"fmt"
	"testing"

	"github.com/gounits/goexcel/parse"
)

func TestXLS(t *testing.T) {
	xls := parse.XLS{}
	load, err := xls.Load("../data/test.xls")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(load)
}
