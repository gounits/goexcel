// @Time  : 2022/7/3 9:11
// @Email: jtyoui@qq.com

package goexcel

import (
	"github.com/gounits/goexcel/parse"
	"github.com/gounits/goexcel/tools"
	"path"
)

type Excel struct {
	parse.Params
}

func New(filepath string) *Excel {
	e := &Excel{}
	e.Filepath = filepath
	// 默认不指定格式，默认启动是文件后缀名自动推导
	switch path.Ext(filepath) {
	case ".xlsx":
		e.Type = parse.XLSX
	case ".xls":
		e.Type = parse.XLS
	case ".csv":
		e.Type = parse.CSV
	case ".tsv":
		e.Type = parse.TSV
	default:
		e.Type = parse.Default
	}
	return e
}

func (e *Excel) Param() *parse.Params {
	return &e.Params
}

func (e *Excel) Load() (bind *Bind, err error) {
	var (
		rows   [][]string
		reader parse.IParse
	)

	if reader, err = e.Type.Reader(e.Params); err != nil {
		return
	}

	if rows, err = reader.Load(); err != nil {
		return
	}

	if len(rows) == 0 || len(rows[0]) == 0 {
		err = tools.EmptyError
		return
	}

	bind = newBind(rows)
	return
}
