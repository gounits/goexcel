package goexcel

import (
	"fmt"
	"reflect"
	"strconv"
)

// 将二维字符串切片转换为指定结构体类型的切片
func convertToStructs[T any](data [][]string) (t []T, err error) {
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

	// 遍历每一行数据
	for idx, row := range rows {
		if len(row) != len(headers) {
			err = fmt.Errorf("第%d行数据列数与表头不匹配", idx+1)
			return
		}
		// 创建一个新的结构体实例
		var item T

		// 获取结构体的反射值
		val := reflect.ValueOf(&item).Elem()

		// 获取结构体的类型
		typ := val.Type()

		// 查找结构体中与表头匹配的字段
		for j := 0; j < typ.NumField(); j++ {
			field := typ.Field(j)

			// 优先使用excel标签匹配，其次使用字段名匹配
			tag := field.Tag.Get("excel")

			if pos, ok := header2idx[tag]; ok {
				// 找到匹配的字段，进行类型转换并赋值
				if err = setFieldValue(val.Field(j), row[pos]); err != nil {
					err = fmt.Errorf("第%d行第%d列赋值错误: %v", idx+1, pos+1, err)
					return
				}
			}
		}

		t = append(t, item)
	}
	return
}

// setFieldValue 根据字段类型设置对应的值
func setFieldValue(field reflect.Value, value string) (err error) {
	if !field.CanSet() {
		err = fmt.Errorf("字段不可设置")
		return
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var val int64
		if val, err = strconv.ParseInt(value, 10, 64); err != nil {
			return
		}
		field.SetInt(val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var val uint64
		if val, err = strconv.ParseUint(value, 10, 64); err != nil {
			return
		}
		field.SetUint(val)
	case reflect.Float32, reflect.Float64:
		var val float64

		if val, err = strconv.ParseFloat(value, 64); err != nil {
			return
		}
		field.SetFloat(val)
	case reflect.Bool:
		var val bool
		if val, err = strconv.ParseBool(value); err != nil {
			return
		}
		field.SetBool(val)
	default:
		err = fmt.Errorf("不支持的字段类型: %s", field.Kind())
	}

	return
}
