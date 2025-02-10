// @Time  : 2022/7/3 9:11
// @Email: jtyoui@qq.com

package parse

import "github.com/gounits/goexcel/tools"

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
	TEXT
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
		result = f.csv(param.Filepath)
	case TSV:
		result = f.tsv(param.Filepath)
	case TEXT:
		result = f.text(param.Filepath)
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

func (f FileType) csv(filepath string) IParse {
	return nil
}

func (f FileType) tsv(filepath string) IParse {
	return nil
}

func (f FileType) text(filepath string) IParse {
	return nil
}
