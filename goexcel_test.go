package goexcel_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gounits/goexcel"
	"github.com/gounits/goexcel/parse"
)

type Test struct {
	Name string `excel:"name"`
	Age  int    `excel:"age"`
	Sex  string `excel:"sex"`
}

func (t *Test) SheetName() string {
	return "test"
}

func TestLoad(t *testing.T) {
	test, err := goexcel.Load[Test, parse.XLSX]("data/test.xlsx", nil)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(test)
}

func TestSave(t *testing.T) {
	test := []Test{{Name: "张三", Age: 17, Sex: "男"}, {Name: "李四", Age: 18, Sex: "女"}}

	if err := goexcel.Save("data/save.xlsx", test); err != nil {
		t.Error(err)
	}

	save, err := goexcel.Load[Test, *parse.XLSX]("data/save.xlsx", nil)
	if err != nil {
		t.Error(err)
		return
	}

	if save[0] != test[0] {
		t.Error(err)
	}

	if err = os.Remove("data/save.xlsx"); err != nil {
		t.Error(err)
	}
}
