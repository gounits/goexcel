package parse

import "golang.org/x/text/encoding"

// Params 可选参数
type Params struct {
	Filepath   string     // 文件路径
	Type       FileType   // 文件类型：枚举类型
	XlsxParams XLSXParams // XLSX 参数配置
	XlsParams  XLSParams  // XLS 参数配置
	CSVParams  CSVParams  // CSV 参数配置
}

func (p *Params) WithType(typ FileType) *Params {
	p.Type = typ
	return p
}

func (p *Params) WithXLSXIndex(index int) *Params {
	p.XlsxParams.Index = index
	return p
}

func (p *Params) WithXLSXSheetName(name string) *Params {
	p.XlsxParams.SheetName = name
	return p
}

func (p *Params) WithXLSIndex(index int) *Params {
	p.XlsParams.Index = index
	return p
}

func (p *Params) WithXLSSheetName(name string) *Params {
	p.XlsParams.SheetName = name
	return p
}

func (p *Params) WithCSVSep(sep rune) *Params {
	p.CSVParams.Sep = sep
	return p
}

func (p *Params) WithCSVFormat(format encoding.Encoding) *Params {
	p.CSVParams.Format = format
	return p
}
