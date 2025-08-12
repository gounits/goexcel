// @Time  : 2022/7/3 9:11
// @Email: jtyoui@qq.com

package parse

import (
	"github.com/xuri/excelize/v2"
)

type XLSX struct {
	SheetName string
	Index     int
}

func (x *XLSX) Load(path string) (rows [][]string, err error) {
	var (
		file  *excelize.File
		sheet = x.SheetName
	)

	if file, err = excelize.OpenFile(path); err != nil {
		return
	}

	defer file.Close()

	if sheet == "" {
		sheet = file.GetSheetName(x.Index)
	}

	rows, err = file.GetRows(sheet)
	return
}
