package goexcel_test

import (
	"fmt"
	"testing"

	"github.com/gounits/goexcel"
	"github.com/gounits/goexcel/parse"
)

type Test struct {
	Name string `excel:"name"`
	Age  int    `excel:"age"`
	Sex  string `excel:"sex"`
}

func TestLoad(t *testing.T) {
	test, err := goexcel.Load[Test, *parse.CSV]("data/test.CSV", nil)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(test)
}
