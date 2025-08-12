package goexcel

import (
	"reflect"

	"github.com/gounits/goexcel/parse"
)

// Load 加载文件数据，绑定到结构体上
func Load[T any, P parse.IParse](filepath string, config func(parse P)) (t []T, err error) {
	var (
		rows  [][]string
		excel P
	)

	// 初始化对象
	if reflect.TypeOf(excel).Kind() == reflect.Ptr {
		typ := reflect.TypeOf(excel).Elem()
		value := reflect.New(typ)
		excel = value.Interface().(P)
	}

	// 设置自定义参数
	if config != nil {
		config(excel)
	}

	// 读取文件返回数据
	if rows, err = excel.Load(filepath); err != nil {
		return
	}

	// 绑定到结构体上
	t, err = convertToStructs[T](rows)
	return
}
