// @Time  : 2022/7/3 9:11
// @Email: jtyoui@qq.com

package goexcel

import (
	"github.com/xuri/excelize/v2"
	"reflect"
	"strings"
)

// LoadExcel 加载Excel文件
func LoadExcel(filePath string, data any) (err error) {
	var (
		rows  [][]string
		file  *excelize.File
		title = map[string]int{}
	)

	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = NotPtrError
		return
	}

	if file, err = excelize.OpenFile(filePath); err != nil {
		return
	}

	defer file.Close()

	rv := reflect.ValueOf(data)
	rv = reflect.Indirect(rv)

	// 默认获取第一个sheet
	sheet := file.GetSheetName(0)
	{
		one := reflect.New(rv.Type().Elem())
		gst := one.MethodByName("SheetName")
		if gst.IsValid() {
			sheet = gst.Call(nil)[0].String()
		}
	}

	if rows, err = file.GetRows(sheet); err != nil {
		return
	}

	if len(rows) == 0 {
		err = EmptyError
		return
	}

	// 获取第一行的标题
	for i, cell := range rows[0] {
		title[cell] = i
	}

	for _, row := range rows[1:] {
		value := reflect.New(rv.Type().Elem())
		value = reflect.Indirect(value)

		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			tags := field.Tag.Get("excel")
			if tags == "" || tags == "-" {
				continue
			}

			tag, split := getSep(tags)

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
					if err = cp(&v, d); err != nil {
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
