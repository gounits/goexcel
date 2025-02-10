package parse

import (
	"bytes"
	csv2 "encoding/csv"
	"github.com/gounits/goexcel/tools"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"io"
	"os"
)

type CSVParams struct {
	Sep    rune              // 文件分隔符
	Format encoding.Encoding // 文件编码格式
}

type csv struct {
	Filepath string
	CSVParams
}

func CheckUTF8(data []byte) error {
	decoder := unicode.UTF8.NewDecoder()
	if newData, err := decoder.Bytes(data); err != nil {
		return err
	} else {
		for i := 0; i < len(newData); i++ {
			if newData[i] != data[i] {
				return tools.UTF8Error
			}
		}
	}
	return nil
}

func (c *csv) Load() (rows [][]string, err error) {
	var (
		file    *os.File
		content []byte
	)

	if file, err = os.Open(c.Filepath); err != nil {
		return
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	if content, err = io.ReadAll(file); err != nil {
		return
	}

	// 判断不是UTF8直接报错误
	if c.Format == nil {
		size := len(content) - 1
		if size > 256 {
			size = 256
		}
		if err = CheckUTF8(content[:size]); err != nil {
			return
		}
	}

	if c.Format != nil {
		if content, err = c.Format.NewDecoder().Bytes(content); err != nil {
			return
		}
	}

	csv := csv2.NewReader(bytes.NewReader(content))
	if c.Sep > 0 {
		csv.Comment = c.Sep
	}
	csv.FieldsPerRecord = -1
	csv.LazyQuotes = true

	// 读取所有行
	rows, err = csv.ReadAll()
	return
}
