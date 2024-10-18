// @Time  : 2022/7/3 9:11
// @Email: jtyoui@qq.com

package parse

// IParse 解析数据的接口
type IParse interface {
	Load() (rows [][]string, err error)
}
