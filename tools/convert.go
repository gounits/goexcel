// @Time  : 2022/6/21 11:04
// @Email: jtyoui@qq.com
// @Author: ZhangWei

package tools

import (
	"fmt"
	"strconv"
)

type TypeConvert interface {
	~string | ~bool | ~int64 | ~uint64 | ~float64
}

// To 将字符串转为其它格式
func To[T TypeConvert](data string) (value T, err error) {
	if data == "" {
		return
	}

	var flag any

	switch formatType := any(value).(type) {
	case int64:
		flag, err = strconv.ParseInt(data, 10, 64)
	case uint64:
		flag, err = strconv.ParseUint(data, 10, 64)
	case float64:
		flag, err = strconv.ParseFloat(data, 64)
	case bool:
		flag, err = strconv.ParseBool(data)
	case string:
		flag = data
	default:
		return value, fmt.Errorf("不支持 %T 类型", formatType)
	}

	if err != nil {
		return value, fmt.Errorf("转换失败: %w", err)
	}

	value = flag.(T)
	return
}
