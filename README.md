# GoExcel

Excel reading and writing based on Go structs

## Install

     go get github.com/gounits/goexcel

## usage

### 1. Save Excel

```go
package main

import (
	"github.com/gounits/goexcel"
)

type Test struct {
	Name     string   `excel:"name"`
	Age      int      `excel:"age"`
	Sex      string   `excel:"sex"`
	UserName []string `excel:"userName;|"`
	High     int      `excel:"-"`
}

func (*Test) SheetName() string {
	return "test"
}

func main() {
	values := []Test{{Name: "张三", Age: 17, Sex: "男"}, {Name: "李四", Age: 18, Sex: "女"}}
	if err := goexcel.SaveExcel("test.xlsx", values); err != nil {
		panic(err)
	}
}
```

### 2. Load Excel

```go
package main

import (
	"fmt"
	"github.com/gounits/goexcel"
)

type Test struct {
	Name     string   `excel:"name"`
	Age      int      `excel:"age"`
	Sex      string   `excel:"sex"`
	UserName []string `excel:"userName;|"`
	High     int      `excel:"-"`
}

func (*Test) SheetName() string {
	return "test"
}

func main() {
	var data []Test
	if err := goexcel.LoadExcel("test.xlsx", &data); err != nil {
		panic(err)
	}
	fmt.Println(data)
}
```
