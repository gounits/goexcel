// @Time  : 2022/7/3 10:12
// @Email: jtyoui@qq.com

package goexcel

import (
	"fmt"
	"github.com/gounits/goexcel/internal"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strings"
)

// SaveExcel 保存到Excel文件中
// filepath: 文件路径
// data: 保存的数据
func SaveExcel(filepath string, data any) (err error) {
	xlsx := excelize.NewFile()

	rv := reflect.ValueOf(data)

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	// 如果不是切片类型
	if rv.Type().Kind() != reflect.Slice {
		err = internal.NoSliceError
		return
	}

	// 获取切片的长度
	length := rv.Len()
	if length == 0 || rv.IsNil() || rv.IsZero() {
		return internal.EmptyError
	}

	// 设置默认的sheet名称
	sheet := "Sheet1"
	{
		// 获取一个元素的类型
		gst := rv.Index(0).MethodByName("SheetName")
		if gst.IsValid() {
			sheet = gst.Call(nil)[0].String()
		}
	}

	var index int
	if index, err = xlsx.NewSheet(sheet); err != nil {
		return
	}

	for i := 0; i < length; i++ {
		elem := rv.Index(i)
		// 如果是指针类型
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		elemType := elem.Type()

		for j := 0; j < elemType.NumField(); j++ {
			field := elemType.Field(j)
			tags := field.Tag.Get("excel")
			if tags == "" || tags == "-" {
				continue
			}

			// get split sep for tag
			tag, split := internal.GetSep(tags)

			m := j % 26
			n := j / 26
			column := fmt.Sprintf("%c", 'A'+m)
			if n >= 1 {
				n--
				column = fmt.Sprintf("%c%s", 'A'+n, column)
			}

			if i == 0 {
				err = xlsx.SetCellValue(sheet, fmt.Sprintf("%s%d", column, i+1), tag)
			}
			if split != "" {
				vs := elem.Field(j).Interface().([]string)
				value := strings.Join(vs, split)
				err = xlsx.SetCellValue(sheet, fmt.Sprintf("%s%d", column, i+2), value)
			} else {
				err = xlsx.SetCellValue(sheet, fmt.Sprintf("%s%d", column, i+2), elem.Field(j).Interface())
			}
		}
	}
	xlsx.SetActiveSheet(index)
	if sheet != "Sheet1" {
		_ = xlsx.DeleteSheet("Sheet1")
	}
	err = xlsx.SaveAs(filepath)
	return
}
