package parse

import (
	"reflect"
)

/*
IParse 解析数据的接口

输入 filepath 文件的路径，获取文件内容。
内容按文件行读取，一行一列 rows
*/
type IParse interface {
	Load(filepath string) (rows [][]string, err error)
}

// ISheetName 实现该接口可以自定义加载SheetName表
type ISheetName interface {
	SheetName() string
}

// SheetName 输入一个对象object判断是否实现了 ISheetName 接口，如果实现了获取结果
func SheetName(object any) (name string, err error) {
	rv := reflect.ValueOf(object)

	if rv.Type().Kind() == reflect.Ptr && rv.IsNil() {
		typ := reflect.TypeOf(object).Elem()
		rv = reflect.New(typ)
	}

	if rv.Kind() == reflect.Ptr {
		for rv.Elem().Kind() == reflect.Ptr {
			rv = rv.Elem()
		}
	} else if rv.Kind() == reflect.Struct {
		typ := reflect.TypeOf(object)
		rv = reflect.New(typ)
	}

	method := rv.MethodByName("SheetName")

	if method.IsValid() {
		name = method.Call(nil)[0].String()
	}

	return
}
