// @Time  : 2022/7/3 9:11
// @Email: jtyoui@qq.com

package parse

import "github.com/extrame/xls"

type XLS struct {
	SheetName string
	Filepath  string
	Index     int
}

func (x *XLS) Load() (res [][]string, err error) {
	var xlsFile *xls.WorkBook

	if xlsFile, err = xls.Open(x.Filepath, "utf-8"); err != nil {
		return
	}

	var sheet *xls.WorkSheet

	sheet = xlsFile.GetSheet(x.Index)

	if x.SheetName != "" {
		for i := 0; i < xlsFile.NumSheets(); i++ {
			temp := xlsFile.GetSheet(i)
			if temp.Name == x.SheetName {
				sheet = temp
				break
			}
		}
	}

	if sheet.MaxRow != 0 {
		sheet.MaxRow += 1
		temp := make([][]string, sheet.MaxRow)
		for i := 0; i < int(sheet.MaxRow); i++ {
			row := sheet.Row(i)
			data := make([]string, 0)
			if row.LastCol() > 0 {
				for j := 0; j < row.LastCol(); j++ {
					col := row.Col(j)
					data = append(data, col)
				}
				temp[i] = data
			}
		}
		res = append(res, temp...)
	}
	return
}
