// @Time  : 2022/7/3 10:12
// @Email: jtyoui@qq.com

package goexcel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strings"
)

// SaveExcel save data to excel.
//
// must be implemented Excel interface.
func SaveExcel[T ~[]E, E IExcel](filepath string, data T) (err error) {
	xlsx := excelize.NewFile()
	if len(data) == 0 {
		return fmt.Errorf("data is empty")
	}
	sheet := data[0].GetSheetName()
	index := xlsx.NewSheet(sheet)

	s := reflect.ValueOf(data)

	for i := 0; i < s.Len(); i++ {
		elem := s.Index(i)
		// drop ptr
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
			tag, split := getSep(tags)

			column := 'A' + j
			if i == 0 {
				err = xlsx.SetCellValue(sheet, fmt.Sprintf("%c%d", column, i+1), tag)
			}
			if split != "" {
				vs := elem.Field(j).Interface().([]string)
				value := strings.Join(vs, split)
				err = xlsx.SetCellValue(sheet, fmt.Sprintf("%c%d", column, i+2), value)
			} else {
				err = xlsx.SetCellValue(sheet, fmt.Sprintf("%c%d", column, i+2), elem.Field(j).Interface())
			}
		}
	}
	xlsx.SetActiveSheet(index)
	xlsx.DeleteSheet("Sheet1")
	err = xlsx.SaveAs(filepath)
	return
}
