package goexcel

import (
	"errors"
	"reflect"

	"github.com/gounits/goexcel/internal"
	"github.com/gounits/goexcel/parse"
	"github.com/xuri/excelize/v2"
)

// Load 加载文件数据，绑定到结构体上
func Load[T any, P parse.IParse](filepath string, config func(parse P)) (t []T, err error) {
	var (
		t1        T
		sheetName string
		rows      [][]string
		excel     P
	)

	// 判断泛型有没有实现 ISheetName 接口
	if sheetName, err = internal.SheetName(t1); err != nil {
		return
	}

	// 初始化对象
	if reflect.TypeOf(excel).Kind() == reflect.Ptr {
		typ := reflect.TypeOf(excel).Elem()
		value := reflect.New(typ)
		if sheetName != "" {
			field := value.Elem().FieldByName("SheetName")
			if field.IsValid() && field.CanSet() {
				field.SetString(sheetName)
			}
		}
		excel = value.Interface().(P)
	}

	// 设置自定义参数
	if config != nil {
		config(excel)
	}

	// 读取文件返回数据
	if rows, err = excel.Load(filepath); err != nil {
		return
	}

	// 绑定到结构体上
	t, err = internal.ConvertToStructs[T](rows)
	return
}

func Save[T any](filepath string, data []T) (err error) {
	var (
		t     T
		sheet string
		index int
	)

	xlsx := excelize.NewFile()

	defer xlsx.Close()

	rv := reflect.ValueOf(data)

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	// 如果不是切片类型
	if rv.Type().Kind() != reflect.Slice {
		err = errors.New("必须输入切片类型")
		return
	}

	// 获取切片的长度
	length := rv.Len()
	{
		if length == 0 || rv.IsNil() || rv.IsZero() {
			err = errors.New("数据为空！")
			return
		}

		if length >= 104_8576 {
			err = errors.New("错误Excel2007(xlsx)以后版本最大行列是1048576行16384列")
			return
		}
	}

	// 设置默认的 sheet 名称
	if sheet, err = internal.SheetName(t); err != nil {
		return
	}

	if sheet == "" {
		sheet = "Sheet1"
	}

	if index, err = xlsx.NewSheet(sheet); err != nil {
		return
	}

	for i := 0; i < length; i++ {
		elem := rv.Index(i)

		// 如果是指针类型
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		typ := elem.Type()

		for j := 0; j < typ.NumField(); j++ {
			var cell, tag string

			field := typ.Field(j)

			if tag = field.Tag.Get("excel"); tag == "" || tag == "-" {
				continue
			}

			// 设置表头
			if i == 0 {
				cell, _ = excelize.CoordinatesToCellName(j+1, 1)
				err = xlsx.SetCellValue(sheet, cell, tag)
			}

			if cell, err = excelize.CoordinatesToCellName(j+1, i+2); err != nil {
				return
			}

			err = xlsx.SetCellValue(sheet, cell, elem.Field(j).Interface())
		}
	}

	if xlsx.SetActiveSheet(index); sheet != "Sheet1" {
		if err = xlsx.DeleteSheet("Sheet1"); err != nil {
			return
		}
	}
	err = xlsx.SaveAs(filepath)
	return
}
