package parse

import (
	"encoding/csv"
	"os"
)

type CSV struct {
	Comment rune
}

func (c CSV) Load(path string) (rows [][]string, err error) {
	var file *os.File

	if file, err = os.Open(path); err != nil {
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	if c.Comment > 0 {
		reader.Comma = c.Comment
	}
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true
	rows, err = reader.ReadAll()
	return
}
