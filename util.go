package linuxtool

import "strconv"

func ParseInt(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func ParseInt32(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 32)
}

func ParseInt64(s string) int64 {
	i64, _ := strconv.ParseInt(s, 10, 64)
	return i64
}

func ParseIntValue(s string) int {
	i64, _ := strconv.ParseInt(s, 10, strconv.IntSize)
	return int(i64)
}

func ParseUint(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func ParseUint32(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 32)
}
func ParseUint64(s string) uint64 {
	i64, _ := strconv.ParseUint(s, 10, 64)
	return i64
}
func ParseHexUint(s string) (uint64, error) {
	return strconv.ParseUint(s, 16, 64)
}

func ParseHexUint64(s string) uint64 {
	i64, _ := strconv.ParseUint(s, 16, 64)
	return i64
}

func ParseHexByte(s string) byte {
	i64, _ := strconv.ParseUint(s, 16, 8)
	return byte(i64)
}

func ParseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func ParseFloat64(s string) float64 {
	f64, _ := strconv.ParseFloat(s, 64)
	return f64
}
