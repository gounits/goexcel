// @Time  : 2022/7/3 9:11
// @Email: jtyoui@qq.com

package goexcel

import (
	"github.com/xuri/excelize/v2"
	"reflect"
	"strings"
)

// LoadExcel loading Excel file
//
// must be implemented Excel interface.
func LoadExcel[T IExcel](filePath string) (data []T, err error) {
	// open Excel file
	var file *excelize.File
	if file, err = excelize.OpenFile(filePath); err != nil {
		return
	}

	defer func() {
		// Close the spreadsheet.
		if err = file.Close(); err != nil {
			return
		}
	}()

	Object := reflect.TypeOf(data).Elem()
	if Object.Kind() == reflect.Ptr {
		Object = Object.Elem()
	}

	// get sheet name from Excel interface
	sheetName := reflect.New(Object).Interface().(T).GetSheetName()

	rows, err := file.GetRows(sheetName)
	if err != nil {
		return
	}

	data = make([]T, len(rows)-1)

	// get title row
	title := map[string]int{}
	for i, cell := range rows[0] {
		title[cell] = i
	}

	// get data row
	for index, row := range rows[1:] {
		t := reflect.New(Object).Interface().(T)

		value := reflect.ValueOf(&t)
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		value = reflect.Indirect(value)

		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			tags := field.Tag.Get("excel")
			if tags == "" || tags == "-" {
				continue
			}

			// get split sep for tag
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

		data[index] = t
	}
	return
}
