// @Time  : 2022/7/3 9:11
// @Email: jtyoui@qq.com

package parse

import (
	"errors"

	"github.com/extrame/xls"
)

type XLS struct {
	SheetName string
	Index     int
}

func (x XLS) Load(path string) (res [][]string, err error) {
	var (
		file  *xls.WorkBook
		sheet *xls.WorkSheet
	)

	if file, err = xls.Open(path, "utf-8"); err != nil {
		return
	}

	if sheet = file.GetSheet(x.Index); x.SheetName != "" {
		for i := 0; i < file.NumSheets(); i++ {
			if temp := file.GetSheet(i); temp != nil && temp.Name == x.SheetName {
				sheet = temp
				break
			}
		}
	}

	if sheet == nil {
		err = errors.New("获取XLS Sheet 为空")
		return
	}

	if sheet.MaxRow == 0 {
		return
	}

	sheet.MaxRow += 1

	for i := 0; i < int(sheet.MaxRow); i++ {
		row := sheet.Row(i)
		data := make([]string, row.LastCol())

		for j := 0; j < row.LastCol(); j++ {
			data[i] = row.Col(j)
		}

		res = append(res, data)
	}

	return
}
