// @Time  : 2022/8/30 14:06
// @Email: jtyoui@qq.com

package tools

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	EmptyError   = errors.New("数据为空")
	NoSliceError = errors.New("数据不是切片类型")
	NotPtrError  = errors.New("数据必须是指针类型")
	TypeError    = errors.New("不支持该文件类型")
)

// ISheet 保存到Excel中的sheet名称
type ISheet interface {
	SheetName() string
}

// isInt 检查是否为有符号整数类型
func isInt(k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}

// isUint 检查是否为无符号整数类型
func isUint(k reflect.Kind) bool {
	switch k {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	default:
		return false
	}
}

// isFloat 检查是否为浮点数类型
func isFloat(k reflect.Kind) bool {
	switch k {
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// Copy 将字符串转换成对应的类型
func Copy(v1 *reflect.Value, b string) (err error) {
	switch t := v1.Kind(); {
	case t == reflect.Bool:
		var value bool
		if value, err = To[bool](b); err != nil {
			return
		}
		v1.SetBool(value)
	case isInt(t):
		var value int64
		if value, err = To[int64](b); err != nil {
			return
		}
		v1.SetInt(value)
	case isUint(t):
		var value uint64
		if value, err = To[uint64](b); err != nil {
			return
		}
		v1.SetUint(value)
	case isFloat(t):
		var value float64
		if value, err = To[float64](b); err != nil {
			return
		}
		v1.SetFloat(value)
	case t == reflect.String:
		v1.SetString(b)
	default:
		return fmt.Errorf("不支持 %T 格式", t.String())
	}
	return
}

// GetSep 根据tag获取分隔符
func GetSep(tags string) (tag string, split string) {
	if strings.Contains(tags, ";") {
		sep := strings.SplitN(tags, ";", 2)
		split = strings.TrimSpace(strings.ToLower(sep[1]))

		// 如果是空格分割，需要用 space 替换
		if split == "space" {
			split = " "
		}

		tag = sep[0]
		return
	}
	tag = tags
	return
}
