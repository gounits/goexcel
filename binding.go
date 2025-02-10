// @Time  : 2022/7/3 9:11
// @Email: jtyoui@qq.com

package goexcel

import (
	"github.com/gounits/goexcel/tools"
	"reflect"
	"strings"
)

type Bind struct {
	Title  []string
	Values [][]string
}

func newBind(value [][]string) *Bind {
	return &Bind{Title: value[0], Values: value[1:]}
}

func (b *Bind) value(data any) (rv reflect.Value, err error) {
	// 必须是指针类型
	if t := reflect.TypeOf(data); t.Kind() != reflect.Ptr {
		err = tools.NotPtrError
		return
	}

	// 获取指针类型的数据类型
	rv = reflect.ValueOf(data)
	rv = reflect.Indirect(rv)
	return
}

func (b *Bind) BindStruct(data any) (err error) {
	var (
		rv    reflect.Value
		title = make(map[string]int)
	)

	// 获取反射类型
	if rv, err = b.value(data); err != nil {
		return
	}

	// 获取第一行的标题
	for i, cell := range b.Title {
		title[cell] = i
	}

	for _, row := range b.Values {
		value := reflect.New(rv.Type().Elem())
		value = reflect.Indirect(value)

		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			tags := field.Tag.Get("excel")
			if tags == "" || tags == "-" {
				continue
			}

			tag, split := tools.GetSep(tags)

			if j, ok := title[tag]; ok {
				v := value.Field(i)

				var d string
				if len(row) > j {
					d = row[j]
				}

				if d == "" {
					continue
				}

				if split == "" {
					if err = tools.Copy(&v, d); err != nil {
						return
					}
				} else {
					vs := strings.Split(row[j], split)
					v.Set(reflect.MakeSlice(v.Type(), len(vs), len(vs)))
					for k, v1 := range vs {
						v.Index(k).SetString(strings.TrimSpace(v1))
					}
				}
			}
		}
		rv.Set(reflect.Append(rv, value))
	}
	return
}
