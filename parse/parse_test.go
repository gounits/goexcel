package parse_test

import (
	"fmt"
	"testing"

	"github.com/gounits/goexcel/parse"
)

type Test struct {
	Name string `excel:"name"`
	Age  int    `excel:"age"`
	Sex  string `excel:"sex"`
}

func (t Test) SheetName() string {
	return "test"
}

func TestSheetName(t *testing.T) {
	t1 := Test{}
	n1, err := parse.SheetName(t1)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(1, n1)

	n2, err := parse.SheetName(&t1)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(2, n2)

	t2 := &Test{}
	n3, err := parse.SheetName(&t2)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(3, n3)

	var t3 Test
	n4, err := parse.SheetName(t3)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(4, n4)

	var t4 *Test
	n5, err := parse.SheetName(t4)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(5, n5)
}
