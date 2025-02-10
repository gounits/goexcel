// @Time  : 2022/7/3 22:17
// @Email: jtyoui@qq.com

package goexcel_test

import (
	"fmt"
	"github.com/gounits/goexcel"
	"github.com/gounits/goexcel/parse"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/unicode"
	"os"
	"testing"
)

type Test struct {
	Name     string   `excel:"name"`
	Age      int      `excel:"age"`
	Sex      string   `excel:"sex"`
	UserName []string `excel:"userName;|"`
	High     int      `excel:"-"`
}

func (Test) SheetName() string {
	return "test"
}

func TestSaveExcel(t *testing.T) {
	values := []Test{{Name: "张三", Age: 17, Sex: "男", UserName: []string{"a", "b"}}, {Name: "李四", Age: 18, Sex: "女"}}
	err := goexcel.SaveExcel("test.xlsx", values)
	assert.NoError(t, err)

	var (
		test []Test
		bind *goexcel.Bind
	)

	excel := goexcel.New("test.xlsx")
	excel.WithXLSXSheetName("test").WithType(parse.XLSX)

	if bind, err = excel.Load(); err != nil {
		assert.NoError(t, err)
		return
	}

	if err = bind.BindStruct(&test); err != nil {
		assert.NoError(t, err)
		return
	}

	assert.Equal(t, test, values)

	_ = os.Remove("test.xlsx")
}

func ExampleSaveExcel() {
	/***
	type Test struct {
		Name     string   `excel:"name"`
		Age      int      `excel:"age"`
		Sex      string   `excel:"sex"`
		UserName []string `excel:"userName;|"`
		High     int      `excel:"-"`
	}

	func (Test) SheetName() string {
		return "test"
	}

	*/

	values := []Test{{Name: "张三", Age: 17, Sex: "男"}, {Name: "李四", Age: 18, Sex: "女", UserName: []string{"张伟", "廖小沫"}}}
	/***
	name	age	sex
	张三		17	男
	李四		18	女
	*/
	err := goexcel.SaveExcel("test.xlsx", values)
	_ = os.Remove("test.xlsx")
	fmt.Println(err)
	// Output:
	// <nil>
}

func ExampleExcel_Load() {
	/***
	type Test struct {
		Name     string   `excel:"name"`
		Age      int      `excel:"age"`
		Sex      string   `excel:"sex"`
		UserName []string `excel:"userName;|"`
		High     int      `excel:"-"`
	}

	func (Test) SheetName() string {
		return "test"
	}

	*/

	values := []*Test{{Name: "张三", Age: 17, Sex: "男"}, {Name: "李四", Age: 18, Sex: "女"}}
	/***
	name	age	sex
	张三		17	男
	李四		18	女
	*/

	var (
		test []Test
		bind *goexcel.Bind
		err  error
	)

	if err = goexcel.SaveExcel("test.xlsx", values); err != nil {
		panic(err)
	}

	if bind, err = goexcel.New("test.xlsx").Load(); err != nil {
		panic(err)
	}

	if err = bind.BindStruct(&test); err != nil {
		panic(err)
	}

	fmt.Println(test)
	_ = os.Remove("test.xlsx")
	// Output:
	// [{张三 17 男 [] 0} {李四 18 女 [] 0}]
}

func ExampleCSV_Load() {
	var (
		test []Test
		bind *goexcel.Bind
		err  error
	)

	excel := goexcel.New("test.csv")
	excel.WithCSVFormat(unicode.UTF8BOM)
	if bind, err = excel.Load(); err != nil {
		panic(err)
	}

	if err = bind.BindStruct(&test); err != nil {
		panic(err)
	}

	fmt.Println(test)
	// Output:
	// [{张三 17 男 [] 0} {李四 18 女 [] 0}]
}
