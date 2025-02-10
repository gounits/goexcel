// @Time  : 2022/7/3 9:11
// @Email: jtyoui@qq.com

package parse

import (
	"github.com/gounits/goexcel/tools"
)

// IParse 解析数据的接口
type IParse interface {
	Load() (rows [][]string, err error)
}

type FileType int

const (
	XLSX FileType = iota
	XLS
	CSV
	TSV
	Default
)

func (f FileType) Reader(param Params) (result IParse, err error) {
	switch f {
	case Default:
		err = tools.TypeError
	case XLSX:
		result = f.xlsx(param.Filepath, param.XlsxParams)
	case XLS:
		result = f.xls(param.Filepath, param.XlsParams)
	case CSV:
		result = f.csv(param.Filepath, param.CSVParams)
	case TSV:
		param.CSVParams.Sep = '\t'
		result = f.csv(param.Filepath, param.CSVParams)
	default:
		err = tools.TypeError
	}
	return
}

func (f FileType) xlsx(filepath string, params XLSXParams) IParse {
	return &xlsx{filepath, params}
}

func (f FileType) xls(filepath string, params XLSParams) IParse {
	return &xls{filepath, params}
}

func (f FileType) csv(filepath string, params CSVParams) IParse {
	return &csv{filepath, params}
}
