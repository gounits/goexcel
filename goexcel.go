// @Time  : 2022/7/3 9:11
// @Email: jtyoui@qq.com

package goexcel

import (
	"github.com/gounits/goexcel/internal"
	"github.com/gounits/goexcel/parse"
	"reflect"
	"strings"
)

type option struct {
	filepath  string
	index     int
	sheetName string
}

type Excel struct {
	option
	value reflect.Value
}

func New(filepath string) *Excel {
	e := &Excel{}
	e.filepath = filepath
	return e
}

func (e *Excel) WithIndex(index int) *Excel {
	e.index = index
	return e
}

func (e *Excel) WithSheetName(name string) *Excel {
	e.sheetName = name
	return e
}

func (e *Excel) parse() (p parse.IParse) {
	if strings.HasSuffix(e.filepath, ".xlsx") {
		p = &parse.XLSX{
			Index:     e.index,
			SheetName: e.sheetName,
			Filepath:  e.filepath,
		}
	}

	if strings.HasSuffix(e.filepath, ".xls") {
		p = &parse.XLS{
			Index:     e.index,
			SheetName: e.sheetName,
			Filepath:  e.filepath,
		}
	}

	return
}

func (e *Excel) Load() (bind *Bind, err error) {
	var rows [][]string

	if rows, err = e.parse().Load(); err != nil {
		return
	}

	if len(rows) == 0 || len(rows[0]) == 0 {
		err = internal.EmptyError
		return
	}

	bind = newBind(rows)
	return
}
