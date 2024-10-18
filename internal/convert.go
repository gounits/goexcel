// @Time  : 2022/6/21 11:04
// @Email: jtyoui@qq.com
// @Author: ZhangWei

package internal

import (
	"reflect"
	"strconv"
)

type TypeConvert interface {
	~string | ~bool | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// To  string convert to f.TypeConvert type
func To[T TypeConvert](data string) (value T, err error) {
	// get value type
	valueType := reflect.TypeOf(value)

	if data == "" {
		return
	}
	var flag any

	switch valueType.Kind() {
	case reflect.Int:
		flag, err = strconv.Atoi(data)
	case reflect.Int8:
		var parse int64
		parse, err = strconv.ParseInt(data, 10, 8)
		flag = int8(parse)
	case reflect.Int16:
		var parse int64
		parse, err = strconv.ParseInt(data, 10, 16)
		flag = int16(parse)
	case reflect.Int32:
		var parse int64
		parse, err = strconv.ParseInt(data, 10, 32)
		flag = int32(parse)
	case reflect.Int64:
		flag, err = strconv.ParseInt(data, 10, 64)
	case reflect.Uint:
		var parse uint64
		parse, err = strconv.ParseUint(data, 10, 0)
		flag = uint(parse)
	case reflect.Uint8:
		var parse uint64
		parse, err = strconv.ParseUint(data, 10, 8)
		flag = uint8(parse)
	case reflect.Uint16:
		var parse uint64
		parse, err = strconv.ParseUint(data, 10, 16)
		flag = uint16(parse)
	case reflect.Uint32:
		var parse uint64
		parse, err = strconv.ParseUint(data, 10, 32)
		flag = uint32(parse)
	case reflect.Uint64:
		flag, err = strconv.ParseUint(data, 10, 64)
	case reflect.Float32:
		var parse float64
		parse, err = strconv.ParseFloat(data, 32)
		flag = float32(parse)
	case reflect.Float64:
		flag, err = strconv.ParseFloat(data, 64)
	case reflect.Bool:
		flag, err = strconv.ParseBool(data)
	case reflect.String:
		flag = data
	default:
		flag = data
	}

	if err == nil {
		value = flag.(T)
	}

	return
}
