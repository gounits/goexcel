// @Time  : 2022/7/3 22:17
// @Email: jtyoui@qq.com

package goexcel_test

import (
	"fmt"
	"github.com/gounits/goexcel"
	"github.com/stretchr/testify/assert"
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

func (*Test) GetSheetName() string {
	return "test"
}

func TestSaveExcel(t *testing.T) {
	values := []*Test{{Name: "张三", Age: 17, Sex: "男", UserName: []string{"a", "b"}}, {Name: "李四", Age: 18, Sex: "女"}}
	err := goexcel.SaveExcel("test.xlsx", values)
	assert.NoError(t, err)

	data, _ := goexcel.LoadExcel[*Test]("test.xlsx")
	assert.Equal(t, data, values)

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

	func (Test) GetSheetName() string {
		return "test"
	}

	*/

	values := []*Test{{Name: "张三", Age: 17, Sex: "男"}, {Name: "李四", Age: 18, Sex: "女"}}
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

func ExampleLoadExcel() {
	/***
	type Test struct {
		Name     string   `excel:"name"`
		Age      int      `excel:"age"`
		Sex      string   `excel:"sex"`
		UserName []string `excel:"userName;|"`
		High     int      `excel:"-"`
	}

	func (Test) GetSheetName() string {
		return "test"
	}

	*/

	values := []*Test{{Name: "张三", Age: 17, Sex: "男"}, {Name: "李四", Age: 18, Sex: "女"}}
	/***
	name	age	sex
	张三		17	男
	李四		18	女
	*/
	_ = goexcel.SaveExcel("test.xlsx", values)
	data, _ := goexcel.LoadExcel[*Test]("test.xlsx")
	fmt.Println(data[0])
	_ = os.Remove("test.xlsx")
	// Output:
	// &{张三 17 男 [] 0}
}
