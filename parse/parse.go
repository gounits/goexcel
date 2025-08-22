package parse

/*
IParse 解析数据的接口

输入 filepath 文件的路径，获取文件内容。
内容按文件行读取，一行一列 rows
*/
type IParse interface {
	Load(filepath string) (rows [][]string, err error)
}

// ISheetName 实现该接口可以自定义加载SheetName表
type ISheetName interface {
	SheetName() string
}
