// @Time  : 2022/6/21 11:12
// @Email: jtyoui@qq.com
// @Author: ZhangWei

package tools_test

import (
	"github.com/gounits/goexcel/tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTo(t *testing.T) {
	s5, _ := tools.To[int64]("65")
	assert.Equal(t, s5, int64(65))

	s10, _ := tools.To[uint64]("65")
	assert.Equal(t, s10, uint64(65))

	s12, _ := tools.To[float64]("65")
	assert.Equal(t, s12, float64(65))

	s13, _ := tools.To[string]("65")
	assert.Equal(t, s13, "65")

	s14, _ := tools.To[bool]("true")
	assert.Equal(t, s14, true)

	s15, _ := tools.To[bool]("false")
	assert.Equal(t, s15, false)

	s16, _ := tools.To[bool]("1")
	assert.Equal(t, s16, true)

	s17, _ := tools.To[bool]("0")
	assert.Equal(t, s17, false)
}
