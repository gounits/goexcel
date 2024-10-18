// @Time  : 2022/7/3 9:11
// @Email: jtyoui@qq.com

package parse

import (
	"github.com/xuri/excelize/v2"
	"log"
)

type XLSX struct {
	Filepath  string
	SheetName string
	Index     int
}

func (X *XLSX) Load() (rows [][]string, err error) {
	var (
		file  *excelize.File
		sheet = X.SheetName
	)

	if file, err = excelize.OpenFile(X.Filepath); err != nil {
		return
	}

	defer func(file *excelize.File) {
		if err := file.Close(); err != nil {
			log.Printf("关闭文件失败: %s", err.Error())
		}
	}(file)

	// 默认获取第一个sheet
	if sheet == "" {
		sheet = file.GetSheetName(X.Index)
	}

	if rows, err = file.GetRows(sheet); err != nil {
		return
	}

	return
}
