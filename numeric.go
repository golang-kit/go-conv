package conv

import (
	"strconv"
	"time"
)

// Do not edit this file by hand as it is generated using generate_test.go from
// testdata/numeric.go.tpl. To make changes edit the .tpl file and type run
// `go generate` to update the file. This process may be removed when conversions
// are solidified.

// NumericConverter interface groups conversion of Numeric Types as described
// in the Go language spec. For more information:
//   https://golang.org/ref/spec#Numeric_types
type NumericConverter interface {
	Complex64Converter
	Complex128Converter
	Float32Converter
	Float64Converter
	IntConverter
	Int8Converter
	Int16Converter
	Int32Converter
	Int64Converter
	UintConverter
	Uint8Converter
	Uint16Converter
	Uint32Converter
	Uint64Converter
}

// Complex64Converter interface allows a value to be converted to a complex64.
type Complex64Converter interface {
	Complex64(from interface{}) complex64
}

// Complex128Converter interface allows a value to be converted to a complex128.
type Complex128Converter interface {
	Complex128(from interface{}) complex128
}

// Float32Converter interface allows a value to be converted to a float32.
type Float32Converter interface {
	Float32(from interface{}) float32
}

// Float64Converter interface allows a value to be converted to a float64.
type Float64Converter interface {
	Float64(from interface{}) float64
}

// IntConverter interface allows a value to be converted to a int.
type IntConverter interface {
	Int(from interface{}) int
}

// Int8Converter interface allows a value to be converted to a int8.
type Int8Converter interface {
	Int8(from interface{}) int8
}

// Int16Converter interface allows a value to be converted to a int16.
type Int16Converter interface {
	Int16(from interface{}) int16
}

// Int32Converter interface allows a value to be converted to a int32.
type Int32Converter interface {
	Int32(from interface{}) int32
}

// Int64Converter interface allows a value to be converted to a int64.
type Int64Converter interface {
	Int64(from interface{}) int64
}

// UintConverter interface allows a value to be converted to a uint.
type UintConverter interface {
	Uint(from interface{}) uint
}

// Uint8Converter interface allows a value to be converted to a uint8.
type Uint8Converter interface {
	Uint8(from interface{}) uint8
}

// Uint16Converter interface allows a value to be converted to a uint16.
type Uint16Converter interface {
	Uint16(from interface{}) uint16
}

// Uint32Converter interface allows a value to be converted to a uint32.
type Uint32Converter interface {
	Uint32(from interface{}) uint32
}

// Uint64Converter interface allows a value to be converted to a uint64.
type Uint64Converter interface {
	Uint64(from interface{}) uint64
}

// Complex64 converts the given value to a complex64.
func Complex64(v interface{}) complex64 {
	v = Indirect(v)

	switch T := v.(type) {
	case Complex64Converter:
		if T != nil {
			return T.Complex64(T)
		}
	case complex64:
		return T
	}
	return complex(Float32(v), float32(0))
}

// Complex128 converts the given value to a complex128.
func Complex128(v interface{}) complex128 {
	v = Indirect(v)

	switch T := v.(type) {
	case Complex128Converter:
		if T != nil {
			return T.Complex128(T)
		}
	case Complex64Converter:
		if T != nil {
			i := T.Complex64(T)
			return complex(float64(real(i)), float64(imag(i)))
		}
	case complex128:
		return T
	}
	return complex(Float64(v), float64(0))
}

// Float64 converts the given value to a float64.
func Float64(v interface{}) float64 {
	v = Indirect(v)

	switch T := v.(type) {
	case Float64Converter:
		if T != nil {
			return T.Float64(T)
		}
	case bool:
		if T != false {
			return 1
		}
	case string:
		if parsed, err := strconv.ParseFloat(T, 64); err == nil {
			return float64(parsed)
		}
		if parsed, err := strconv.ParseInt(T, 10, 0); err == nil {
			return float64(parsed)
		}
		if parsed, err := strconv.ParseBool(T); err == nil {
			if parsed {
				return float64(1)
			}
			return 0
		}
	case time.Duration:
		return float64(T)
	case time.Time:
		return float64(T.Unix())
	case complex64:
		return float64(real(T))
	case complex128:
		return float64(real(T))
	case uint:
		return float64(T)
	case uint8:
		return float64(T)
	case uint16:
		return float64(T)
	case uint32:
		return float64(T)
	case uint64:
		return float64(T)
	case float32:
		return float64(T)
	case float64:
		return T
	case int:
		return float64(T)
	case int8:
		return float64(T)
	case int16:
		return float64(T)
	case int32:
		return float64(T)
	case int64:
		return float64(T)
	}
	return 0
}

// Float32 converts value to float32 by calling Float64 and using the defined Go
// numeric conversion rules as described in the language spec.
func Float32(v interface{}) float32 {
	switch T := v.(type) {
	case Float32Converter:
		if T != nil {
			return T.Float32(T)
		}
	}
	return float32(Float64(v))
}

// Int64 converts the given value to a int64.
func Int64(v interface{}) int64 {
	v = Indirect(v)

	switch T := v.(type) {
	case Int64Converter:
		if T != nil {
			return T.Int64(T)
		}
	case bool:
		if T != false {
			return 1
		}
	case string:
		if parsed, err := strconv.ParseInt(T, 10, 0); err == nil {
			return int64(parsed)
		}
		if parsed, err := strconv.ParseFloat(T, 64); err == nil {
			return int64(parsed)
		}
		if parsed, err := strconv.ParseBool(T); err == nil {
			if parsed {
				return int64(1)
			}
			return 0
		}
	case time.Duration:
		return int64(T)
	case time.Time:
		return int64(T.Unix())
	case complex64:
		return int64(real(T))
	case complex128:
		return int64(real(T))
	case uint:
		return int64(T)
	case uint8:
		return int64(T)
	case uint16:
		return int64(T)
	case uint32:
		return int64(T)
	case uint64:
		return int64(T)
	case float32:
		return int64(T)
	case float64:
		return int64(T)
	case int:
		return int64(T)
	case int8:
		return int64(T)
	case int16:
		return int64(T)
	case int32:
		return int64(T)
	case int64:
		return T
	}
	return 0
}

// Int converts value to int by calling Int64 and using the defined Go
// numeric conversion rules as described in the language spec.
func Int(v interface{}) int {
	switch T := v.(type) {
	case IntConverter:
		if T != nil {
			return T.Int(T)
		}
	}
	return int(Int64(v))
}

// Int8 converts value to int8 by calling Int64 and using the defined Go
// numeric conversion rules as described in the language spec.
func Int8(v interface{}) int8 {
	switch T := v.(type) {
	case Int8Converter:
		if T != nil {
			return T.Int8(T)
		}
	}
	return int8(Int64(v))
}

// Int16 converts value to int16 by calling Int64 and using the defined Go
// numeric conversion rules as described in the language spec.
func Int16(v interface{}) int16 {
	switch T := v.(type) {
	case Int16Converter:
		if T != nil {
			return T.Int16(T)
		}
	}
	return int16(Int64(v))
}

// Int32 converts value to int32 by calling Int64 and using the defined Go
// numeric conversion rules as described in the language spec.
func Int32(v interface{}) int32 {
	switch T := v.(type) {
	case Int32Converter:
		if T != nil {
			return T.Int32(T)
		}
	}
	return int32(Int64(v))
}

// Uint64 converts the given value to a uint64.
func Uint64(v interface{}) uint64 {
	v = Indirect(v)

	switch T := v.(type) {
	case Uint64Converter:
		if T != nil {
			return T.Uint64(T)
		}
	case bool:
		if T != false {
			return 1
		}
	case string:
		if parsed, err := strconv.ParseUint(T, 10, 0); err == nil {
			return uint64(parsed)
		}
		if parsed, err := strconv.ParseFloat(T, 64); err == nil {
			if 1 > parsed {
				return uint64(0)
			}
			return uint64(parsed)
		}
		if parsed, err := strconv.ParseBool(T); err == nil {
			if parsed {
				return uint64(1)
			}
			return 0
		}
	case time.Duration:
		return uint64(T)
	case time.Time:
		return uint64(T.Unix())
	case complex64:
		return uint64(real(T))
	case complex128:
		return uint64(real(T))
	case uint:
		return uint64(T)
	case uint8:
		return uint64(T)
	case uint16:
		return uint64(T)
	case uint32:
		return uint64(T)
	case uint64:
		return T
	case float32:
		return uint64(T)
	case float64:
		return uint64(T)
	case int:
		return uint64(T)
	case int8:
		return uint64(T)
	case int16:
		return uint64(T)
	case int32:
		return uint64(T)
	case int64:
		return uint64(T)
	}
	return 0
}

// Uint converts value to uint by calling Uint64 and using the defined Go
// numeric conversion rules as described in the language spec.
func Uint(v interface{}) uint {
	switch T := v.(type) {
	case UintConverter:
		if T != nil {
			return T.Uint(T)
		}
	}
	return uint(Uint64(v))
}

// Uint8 converts value to uint8 by calling Uint64 and using the defined Go
// numeric conversion rules as described in the language spec.
func Uint8(v interface{}) uint8 {
	switch T := v.(type) {
	case Uint8Converter:
		if T != nil {
			return T.Uint8(T)
		}
	}
	return uint8(Uint64(v))
}

// Uint16 converts value to uint16 by calling Uint64 and using the defined Go
// numeric conversion rules as described in the language spec.
func Uint16(v interface{}) uint16 {
	switch T := v.(type) {
	case Uint16Converter:
		if T != nil {
			return T.Uint16(T)
		}
	}
	return uint16(Uint64(v))
}

// Uint32 converts value to uint32 by calling Uint64 and using the defined Go
// numeric conversion rules as described in the language spec.
func Uint32(v interface{}) uint32 {
	switch T := v.(type) {
	case Uint32Converter:
		if T != nil {
			return T.Uint32(T)
		}
	}
	return uint32(Uint64(v))
}

// Indirect will return the value pointed to by v if it is a convertable type
// and not nil. Types that are not pointers are returned without modification.
func Indirect(v interface{}) interface{} {
	switch T := v.(type) {
	case *interface{}:
		if nil != T {
			return Indirect(*T)
		}
	case **interface{}:
		if nil != T {
			return Indirect(*T)
		}
	case *Value:
		if nil != T {
			return Indirect(*T)
		}
	case **Value:
		if nil != T {
			return Indirect(*T)
		}
	case *time.Duration:
		if nil != T {
			return Indirect(*T)
		}
	case **time.Duration:
		if nil != T {
			return Indirect(*T)
		}
	case *time.Time:
		if nil != T {
			return Indirect(*T)
		}
	case **time.Time:
		if nil != T {
			return Indirect(*T)
		}
	case *bool:
		if nil != T {
			return Indirect(*T)
		}
	case **bool:
		if nil != T {
			return Indirect(*T)
		}
	case *string:
		if nil != T {
			return Indirect(*T)
		}
	case **string:
		if nil != T {
			return Indirect(*T)
		}
	case *[]byte:
		if nil != T {
			return Indirect(*T)
		}
	case **[]byte:
		if nil != T {
			return Indirect(*T)
		}
	case *complex64:
		if nil != T {
			return Indirect(*T)
		}
	case **complex64:
		if nil != T {
			return Indirect(*T)
		}
	case *complex128:
		if nil != T {
			return Indirect(*T)
		}
	case **complex128:
		if nil != T {
			return Indirect(*T)
		}
	case *uint:
		if nil != T {
			return Indirect(*T)
		}
	case **uint:
		if nil != T {
			return Indirect(*T)
		}
	case *uint8:
		if nil != T {
			return Indirect(*T)
		}
	case **uint8:
		if nil != T {
			return Indirect(*T)
		}
	case *uint16:
		if nil != T {
			return Indirect(*T)
		}
	case **uint16:
		if nil != T {
			return Indirect(*T)
		}
	case *uint32:
		if nil != T {
			return Indirect(*T)
		}
	case **uint32:
		if nil != T {
			return Indirect(*T)
		}
	case *uint64:
		if nil != T {
			return Indirect(*T)
		}
	case **uint64:
		if nil != T {
			return Indirect(*T)
		}
	case *float32:
		if nil != T {
			return Indirect(*T)
		}
	case **float32:
		if nil != T {
			return Indirect(*T)
		}
	case *float64:
		if nil != T {
			return Indirect(*T)
		}
	case **float64:
		if nil != T {
			return Indirect(*T)
		}
	case *int:
		if nil != T {
			return Indirect(*T)
		}
	case **int:
		if nil != T {
			return Indirect(*T)
		}
	case *int8:
		if nil != T {
			return Indirect(*T)
		}
	case **int8:
		if nil != T {
			return Indirect(*T)
		}
	case *int16:
		if nil != T {
			return Indirect(*T)
		}
	case **int16:
		if nil != T {
			return Indirect(*T)
		}
	case *int32:
		if nil != T {
			return Indirect(*T)
		}
	case **int32:
		if nil != T {
			return Indirect(*T)
		}
	case *int64:
		if nil != T {
			return Indirect(*T)
		}
	case **int64:
		if nil != T {
			return Indirect(*T)
		}
	}
	return v
}
