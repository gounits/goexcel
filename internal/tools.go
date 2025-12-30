package internal

import (
	"errors"
	"fmt"
	"reflect"
)

func newObject(object any) (rv reflect.Value, err error) {
	rv = reflect.ValueOf(object)

	switch rv.Kind() {
	case reflect.Ptr:
		if rv.IsNil() {
			typ := reflect.TypeOf(object).Elem()
			rv = reflect.New(typ)
		}
		if rv.Elem().Kind() == reflect.Ptr {
			rv = rv.Elem()
		}
	case reflect.Struct:
		typ := reflect.TypeOf(object)
		rv = reflect.New(typ)
	default:
		err = errors.New("格式只支持结构体和结构体指针")
	}

	return
}

// SheetName 输入一个对象object判断是否实现了 ISheetName 接口，如果实现了获取结果
func SheetName(object any) (name string, err error) {
	var rv reflect.Value

	if rv, err = newObject(object); err != nil {
		err = errors.Join(err, errors.New("获取 SheetName 名称失败"))
		return
	}

	if method := rv.MethodByName("SheetName"); method.IsValid() {
		name = method.Call(nil)[0].String()
	}
	return
}

// ConvertToStructs 将二维字符串切片转换为指定结构体类型的切片
func ConvertToStructs[T any](data [][]string) (t []T, err error) {
	var (
		item T
		val  reflect.Value
	)

	if len(data) < 2 {
		err = fmt.Errorf("数据至少需要包含表头和一行数据")
		return
	}

	// 获取表头
	headers := data[0]
	header2idx := make(map[string]int)
	for i, header := range headers {
		header2idx[header] = i
	}

	// 获取数据行
	rows := data[1:]

	// 创建结果切片
	t = make([]T, 0, len(rows))

	// 初始化泛型对象并获取类型
	if val, err = newObject(item); err != nil {
		err = errors.Join(err, errors.New("传入的泛型类型不是结构体或结构体指针"))
		return
	}

	val = val.Elem()
	typ := val.Type()

	// 判断泛型传入的是指针吗？
	ptr := reflect.ValueOf(item).Kind() == reflect.Ptr

	// 遍历每一行数据
	for idx, row := range rows {
		if val = reflect.New(typ); ptr {
			item = val.Interface().(T)
		}
		val = val.Elem()

		// 查找结构体中与表头匹配的字段
		for j := 0; j < typ.NumField(); j++ {
			field := typ.Field(j)

			// 优先使用excel标签匹配，其次使用字段名匹配
			tag := field.Tag.Get("excel")

			if pos, ok := header2idx[tag]; ok {
				// [fix] 修复当数据缺失的情况
				value := NA
				if pos < len(row) {
					value = Value(row[pos])
				}

				if err = setFieldValue(val.Field(j), value); err != nil {
					err = fmt.Errorf("第%d行第%d列赋值错误: %v", idx+1, pos+1, err)
					return
				}
			}
		}

		if !ptr {
			item = val.Interface().(T)
		}

		t = append(t, item)
	}
	return
}

// setFieldValue 根据字段类型设置对应的值
func setFieldValue(field reflect.Value, value Value) (err error) {
	if !field.CanSet() {
		err = fmt.Errorf("字段不可设置")
		return
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(value.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var val int64
		if val, err = value.Int64(); err != nil {
			return
		}
		field.SetInt(val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var val uint64
		if val, err = value.Uint64(); err != nil {
			return
		}
		field.SetUint(val)
	case reflect.Float32, reflect.Float64:
		var val float64

		if val, err = value.Float64(); err != nil {
			return
		}
		field.SetFloat(val)
	case reflect.Bool:
		var val bool
		if val, err = value.Bool(); err != nil {
			return
		}
		field.SetBool(val)
	default:
		err = fmt.Errorf("不支持的字段类型: %s", field.Kind())
	}

	return
}

// CollectError 聚合异常 返回 defer 关闭的异常信息
func CollectError(f func() error, dst *error) {
	if f == nil {
		return
	}

	err := f()

	if err == nil {
		return
	}

	if *dst == nil {
		*dst = err
		return
	}

	*dst = errors.Join(*dst, err)
}
