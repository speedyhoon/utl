package utl

import (
	"strconv"
	"strings"
	"testing"
)

func BenchmarkOrdinal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ordinal(uint(i), false)
	}
}

//func BenchmarkOrdinalTrue(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		ordinal(i, true)
//	}
//}
func BenchmarkOrdinal1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ordinal1(i, false)
	}
}

//func BenchmarkOrdinal1True(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		ordinal1(i, true)
//	}
//}
func BenchmarkOrdinal2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ordinal2(i, false)
	}
}

//func BenchmarkOrdinal2True(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		ordinal2(i, true)
//	}
//}
func BenchmarkOrdinal3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ordinal3(i, false)
	}
}

//func BenchmarkOrdinal3True(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		ordinal3(i, true)
//	}
//}
func BenchmarkOrdinal4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ordinal4(i, false)
	}
}

//func BenchmarkOrdinal4True(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		ordinal4(i, true)
//	}
//}
func BenchmarkOrdinal5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ordinal5(i, false)
	}
}

//func BenchmarkOrdinal5True(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		ordinal5(i, true)
//	}
//}
func BenchmarkOrdinal6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ordinal6(i, false)
	}
}

//func BenchmarkOrdinal6True(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		ordinal6(i, true)
//	}
//}
func BenchmarkOrdinal7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ordinal7(i, false)
	}
}

//func BenchmarkOrdinal7True(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		ordinal7(i, true)
//	}
//}
func BenchmarkOrdinal8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ordinal8(i, false)
	}
}

//func BenchmarkOrdinal8True(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		ordinal8(i, true)
//	}
//}
func BenchmarkOrdinal9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ordinal9(i, false)
	}
}

//func BenchmarkOrdinal9True(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		ordinal9(i, true)
//	}
//}
func BenchmarkOrdinal10(b *testing.B) {
	var i uint64
	for ; i < uint64(b.N); i++ {
		ordinal10(i, false)
	}
}

//func BenchmarkOrdinal10True(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		ordinal10(i, true)
//	}
//}
func BenchmarkOrdinal11(b *testing.B) {
	var i uint64
	for ; i < uint64(b.N); i++ {
		ordinal11(i, false)
	}
}

//func BenchmarkOrdinal11True(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		ordinal11(i, true)
//	}
//}
func BenchmarkOrdinal12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ordinal12(i, false)
	}
}

//func BenchmarkOrdinal12True(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		ordinal12(i, true)
//	}
//}
func BenchmarkOrdinal13(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ordinal13(i, false)
	}
}

//func BenchmarkOrdinal13True(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		ordinal13(i, true)
//	}
//}
func BenchmarkOrdinal14(b *testing.B) {
	var i uint
	for ; i < uint(b.N); i++ {
		ordinal14(i, false)
	}
}

//func BenchmarkOrdinal14True(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		ordinal14(i, true)
//	}
//}

func ordinal1(x int, isEqual bool) string {
	switch x % 100 {
	case 11, 12, 13:
		if isEqual {
			return "=" + strconv.Itoa(x) + "th"
		}
		return strconv.Itoa(x) + "th"
	}

	switch x % 10 {
	case 1:
		if isEqual {
			return "=" + strconv.Itoa(x) + "st"
		}
		return strconv.Itoa(x) + "st"
	case 2:
		if isEqual {
			return "=" + strconv.Itoa(x) + "nd"
		}
		return strconv.Itoa(x) + "nd"
	case 3:
		if isEqual {
			return "=" + strconv.Itoa(x) + "rd"
		}
		return strconv.Itoa(x) + "rd"
	}

	if isEqual {
		return "=" + strconv.Itoa(x) + "th"
	}

	return strconv.Itoa(x) + "th"
}
func ordinal2(x int, isEqual bool) string {
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			if isEqual {
				return "=" + strconv.Itoa(x) + "st"
			}
			return strconv.Itoa(x) + "st"
		}
	case 2:
		if x%100 != 12 {
			if isEqual {
				return "=" + strconv.Itoa(x) + "nd"
			}
			return strconv.Itoa(x) + "nd"
		}
	case 3:
		if x%100 != 13 {
			if isEqual {
				return "=" + strconv.Itoa(x) + "rd"
			}
			return strconv.Itoa(x) + "rd"
		}
	}

	if isEqual {
		return "=" + strconv.Itoa(x) + "th"
	}

	return strconv.Itoa(x) + "th"
}

func ordinal3(x int, isEqual bool) string {
	suffix := "th"

	switch x % 10 {
	case 1:
		if x%100 != 11 {
			suffix = "st"
		}
	case 2:
		if x%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if x%100 != 13 {
			suffix = "rd"
		}
	}

	if isEqual {
		return "=" + strconv.Itoa(x) + suffix
	}

	return strconv.Itoa(x) + suffix
}

func ordinal4(x int, isEqual bool) (suffix string) {
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			suffix = "st"
		} else {
			suffix = "th"
		}
	case 2:
		if x%100 != 12 {
			suffix = "nd"
		} else {
			suffix = "th"
		}
	case 3:
		if x%100 != 13 {
			suffix = "rd"
		} else {
			suffix = "th"
		}
	default:
		suffix = "th"
	}

	if isEqual {
		return "=" + strconv.Itoa(x) + suffix
	}

	return strconv.Itoa(x) + suffix
}

func ordinal5(x int, isEqual bool) (suffix string) {
	if isEqual {
		suffix = "="
	}
	suffix += strconv.Itoa(x)

	switch x % 10 {
	case 1:
		if x%100 != 11 {
			return suffix + "st"
		}
	case 2:
		if x%100 != 12 {
			return suffix + "nd"
		}
	case 3:
		if x%100 != 13 {
			return suffix + "rd"
		}
	}

	//if isEqual {
	//	return "=" + strconv.Itoa(x) + suffix
	//}
	//
	return suffix + "th"
}

func ordinal6(x int, isEqual bool) (suffix string) {
	builder := strings.Builder{}
	if isEqual {
		builder.WriteString("=")
	}
	builder.WriteString(strconv.Itoa(x))

	switch x % 10 {
	case 1:
		if x%100 != 11 {
			builder.WriteString("st")
			return builder.String()
		}
	case 2:
		if x%100 != 12 {
			builder.WriteString("nd")
			return builder.String()
		}
	case 3:
		if x%100 != 13 {
			builder.WriteString("rd")
			return builder.String()
		}
	}

	builder.WriteString("th")
	return builder.String()
}

func ordinal7(x int, isEqual bool) (suffix string) {
	builder := strings.Builder{}
	if isEqual {
		builder.WriteByte(61)
	}
	builder.WriteString(strconv.Itoa(x))

	switch x % 10 {
	case 1:
		if x%100 != 11 {
			builder.Write([]byte("st"))
			return builder.String()
		}
	case 2:
		if x%100 != 12 {
			builder.Write([]byte("nd"))
			return builder.String()
		}
	case 3:
		if x%100 != 13 {
			builder.Write([]byte("rd"))
			return builder.String()
		}
	}

	builder.Write([]byte("th"))
	return builder.String()
}

func ordinal8(x int, isEqual bool) (suffix string) {
	if isEqual {
		suffix = "=" + strconv.Itoa(x)
	} else {
		suffix = strconv.Itoa(x)
	}

	switch x % 10 {
	case 1:
		if x%100 != 11 {
			return suffix + "st"
		}
	case 2:
		if x%100 != 12 {
			return suffix + "nd"
		}
	case 3:
		if x%100 != 13 {
			return suffix + "rd"
		}
	}

	return suffix + "th"
}

func ordinal9(x int, isEqual bool) string {
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			if isEqual {
				return "=" + strconv.Itoa(x) + "st"
			}
			return strconv.Itoa(x) + "st"
		} else {
			if isEqual {
				return "=" + strconv.Itoa(x) + "th"
			}

			return strconv.Itoa(x) + "th"
		}
	case 2:
		if x%100 != 12 {
			if isEqual {
				return "=" + strconv.Itoa(x) + "nd"
			}
			return strconv.Itoa(x) + "nd"
		} else {
			if isEqual {
				return "=" + strconv.Itoa(x) + "th"
			}

			return strconv.Itoa(x) + "th"
		}
	case 3:
		if x%100 != 13 {
			if isEqual {
				return "=" + strconv.Itoa(x) + "rd"
			}
			return strconv.Itoa(x) + "rd"
		} else {
			if isEqual {
				return "=" + strconv.Itoa(x) + "th"
			}

			return strconv.Itoa(x) + "th"
		}
	default:
		if isEqual {
			return "=" + strconv.Itoa(x) + "th"
		}

		return strconv.Itoa(x) + "th"
	}
}

func ordinal10(x uint64, isEqual bool) (suffix string) {
	if isEqual {
		suffix = "=" + strconv.FormatUint(x, 10)
	} else {
		suffix = strconv.FormatUint(x, 10)
	}

	switch x % 10 {
	case 1:
		if x%100 != 11 {
			return suffix + "st"
		}
	case 2:
		if x%100 != 12 {
			return suffix + "nd"
		}
	case 3:
		if x%100 != 13 {
			return suffix + "rd"
		}
	}

	return suffix + "th"
}

func ordinal11(x uint64, isEqual bool) (suffix string) {
	//if isEqual {
	//	suffix = "=" + strconv.FormatUint(x,10)
	//} else {
	suffix = strconv.FormatUint(x, 10)
	//}

	switch x % 10 {
	case 1:
		if x%100 != 11 {
			return suffix + "st"
		}
	case 2:
		if x%100 != 12 {
			return suffix + "nd"
		}
	case 3:
		if x%100 != 13 {
			return suffix + "rd"
		}
	}

	return suffix + "th"
}

func ordinal12(x int, isEqual bool) (suffix string) {
	if isEqual {
		suffix = "=" + strconv.FormatUint(uint64(x), 10)
	} else {
		suffix = strconv.FormatUint(uint64(x), 10)
	}

	switch x % 10 {
	case 1:
		if x%100 != 11 {
			return suffix + "st"
		}
	case 2:
		if x%100 != 12 {
			return suffix + "nd"
		}
	case 3:
		if x%100 != 13 {
			return suffix + "rd"
		}
	}

	return suffix + "th"
}

func ordinal13(x int, isEqual bool) string {
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			if isEqual {
				return "=" + strconv.FormatUint(uint64(x), 10) + "st"
			}
			return strconv.FormatUint(uint64(x), 10) + "st"
		} else {
			if isEqual {
				return "=" + strconv.FormatUint(uint64(x), 10) + "th"
			}

			return strconv.FormatUint(uint64(x), 10) + "th"
		}
	case 2:
		if x%100 != 12 {
			if isEqual {
				return "=" + strconv.FormatUint(uint64(x), 10) + "nd"
			}
			return strconv.FormatUint(uint64(x), 10) + "nd"
		} else {
			if isEqual {
				return "=" + strconv.FormatUint(uint64(x), 10) + "th"
			}

			return strconv.FormatUint(uint64(x), 10) + "th"
		}
	case 3:
		if x%100 != 13 {
			if isEqual {
				return "=" + strconv.FormatUint(uint64(x), 10) + "rd"
			}
			return strconv.FormatUint(uint64(x), 10) + "rd"
		} else {
			if isEqual {
				return "=" + strconv.FormatUint(uint64(x), 10) + "th"
			}

			return strconv.FormatUint(uint64(x), 10) + "th"
		}
	default:
		if isEqual {
			return "=" + strconv.FormatUint(uint64(x), 10) + "th"
		}

		return strconv.FormatUint(uint64(x), 10) + "th"
	}
}

func ordinal14(x uint, isEqual bool) (suffix string) {
	if isEqual {
		suffix = "=" + strconv.FormatUint(uint64(x), 10)
	} else {
		suffix = strconv.FormatUint(uint64(x), 10)
	}

	switch x % 10 {
	case 1:
		if x%100 != 11 {
			return suffix + "st"
		}
	case 2:
		if x%100 != 12 {
			return suffix + "nd"
		}
	case 3:
		if x%100 != 13 {
			return suffix + "rd"
		}
	}

	return suffix + "th"
}
