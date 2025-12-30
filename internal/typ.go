package internal

// 特殊字符转换

import (
	"math"
	"strconv"
)

type Value string

const (
	NAN  Value = "NaN"  // 非数字
	INF  Value = "inf"  // 无穷大
	_INF Value = "-inf" // 负无穷大
	NA   Value = ""     // 缺失值
)

func (v Value) String() string {
	return string(v)
}

func (v Value) Int64() (val int64, err error) {
	switch v {
	case NA, NAN:
		val = 0
	case INF:
		val = math.MaxInt64
	case _INF:
		val = math.MinInt64
	default:
		val, err = strconv.ParseInt(v.String(), 10, 64)
	}
	return
}

func (v Value) Uint64() (val uint64, err error) {
	switch v {
	case NA, NAN:
		val = 0
	case INF:
		val = math.Float64bits(1)
	case _INF:
		val = math.Float64bits(-1)
	default:
		val, err = strconv.ParseUint(v.String(), 10, 64)
	}

	return
}

func (v Value) Float64() (val float64, err error) {
	switch v {
	case NA:
		val = 0
	case INF:
		val = math.Inf(1)
	case _INF:
		val = math.Inf(-1)
	case NAN:
		val = math.NaN()
	default:
		val, err = strconv.ParseFloat(v.String(), 64)
	}
	return
}

func (v Value) Bool() (val bool, err error) {
	val, err = strconv.ParseBool(v.String())
	return
}
