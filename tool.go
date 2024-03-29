// @Time  : 2022/8/30 14:06
// @Email: jtyoui@qq.com

package goexcel

import (
	"errors"
	"github.com/gounits/goexcel/internal"
	"reflect"
	"strings"
)

var (
	EmptyError   = errors.New("data is empty")
	NoSliceError = errors.New("data is not slice")
	NotPtrError  = errors.New("data is not ptr")
)

// ISheet 保存到Excel中的sheet名称
type ISheet interface {
	SheetName() string
}

// 将字符串转换成对应的类型
func cp(v1 *reflect.Value, b string) error {
	switch v1.Kind() {
	case reflect.Bool:
		if v2, err := internal.To[bool](b); err != nil {
			return err
		} else {
			v1.SetBool(v2)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v2, err := internal.To[int64](b); err != nil {
			return err
		} else {
			v1.SetInt(v2)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v2, err := internal.To[uint64](b); err != nil {
			return err
		} else {
			v1.SetUint(v2)
		}
	case reflect.Float32, reflect.Float64:
		if v2, err := internal.To[float64](b); err != nil {
			return err
		} else {
			v1.SetFloat(v2)
		}
	case reflect.String:
		v1.SetString(b)
	default:
		return errors.New("unsupported type: " + v1.Type().String())
	}
	return nil
}

// 根据tag获取分隔符
func getSep(tags string) (tag string, split string) {
	if strings.Contains(tags, ";") {
		sep := strings.SplitN(tags, ";", 2)
		tag = sep[0]
		split = strings.TrimSpace(sep[1])

		// 如果是空格分割，需要用 space 替换
		if split == "space" {
			split = " "
		}
	} else {
		tag = tags
	}
	return
}
