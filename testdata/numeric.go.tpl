package conv

import (
	"strconv"
	"time"
)

// Do not edit this file by hand as it is generated using generate_test.go from
// testdata/numeric.go.tpl. To make changes edit the .tpl file and type run
// `go generate` to update the file. This process may be removed when conversions
// are solidified.

{{ with $types := (args "complex64" "complex128" "float32" "float64" "int" "int8" "int16" "int32" "int64" "uint" "uint8" "uint16" "uint32" "uint64") }}
// NumericConverter interface groups conversion of Numeric Types as described
// in the Go language spec. For more information:
//   https://golang.org/ref/spec#Numeric_types
type NumericConverter interface {
	{{- range $type := . }}
	{{ $type | Title }}Converter
	{{- end }}
}

{{- range $type := . }}
{{ template "converterInterface" (args $type) -}}
{{ end }}
{{ end -}}

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

{{ template "converterFunc" (args "float64" (args "ParseFloat(T, 64)" "ParseInt(T, 10, 0)") "false")}}
{{ template "converterFuncProxy" (args "float32" "Float64") }}

{{ template "converterFunc" (args "int64" (args "ParseInt(T, 10, 0)" "ParseFloat(T, 64)") "false") }}
{{ template "converterFuncProxy" (args "int" "Int64") }}
{{ template "converterFuncProxy" (args "int8" "Int64") }}
{{ template "converterFuncProxy" (args "int16" "Int64") }}
{{ template "converterFuncProxy" (args "int32" "Int64") }}

{{ template "converterFunc" (args "uint64" (args "ParseUint(T, 10, 0)" "ParseFloat(T, 64)") "true") }}
{{ template "converterFuncProxy" (args "uint" "Uint64") }}
{{ template "converterFuncProxy" (args "uint8" "Uint64") }}
{{ template "converterFuncProxy" (args "uint16" "Uint64") }}
{{ template "converterFuncProxy" (args "uint32" "Uint64") }}

// Indirect will return the value pointed to by v if it is a convertable type
// and not nil. Types that are not pointers are returned without modification.
func Indirect(v interface{}) interface{} {
	switch T := v.(type) {
	{{- range $from := (args "interface{}" "Value" "time.Duration" "time.Time" ) }}
	{{ template "indirectFuncCase" (args $from ) }}
	{{- end }}
	{{- range $from := (args "bool" "string" "[]byte" "complex64" "complex128" ) }}
	{{ template "indirectFuncCase" (args $from ) }}
	{{- end }}
	{{- range $from := (args "uint" "uint8" "uint16" "uint32" "uint64" "float32") }}
	{{ template "indirectFuncCase" (args $from ) }}
	{{- end }}
	{{- range $from := (args "float64" "int" "int8" "int16" "int32" "int64") }}
	{{ template "indirectFuncCase" (args $from ) }}
	{{- end }}
	}
	return v
}

{{ define "indirectFuncCase" }}
{{- $from := index . 0 -}}
	case *{{ $from }}:
		if nil != T {
			return Indirect(*T)
		}
	case **{{ $from }}:
		if nil != T {
			return Indirect(*T)
		}
{{- end }}

{{ define "converterFuncProxy" }}
{{- $to := index . 0 -}}
{{- $proxyFunc := index . 1 -}}
{{- $thisFunc := Title $to -}}
// {{$thisFunc}} converts value to {{$to}} by calling {{ $proxyFunc }} and using the defined Go
// numeric conversion rules as described in the language spec.
func {{$thisFunc}}(v interface{}) {{$to}} {
	switch T := v.(type) {
	case {{$thisFunc}}Converter:
		if T != nil {
			return T.{{$thisFunc}}(T)
		}
	}
	return {{ printf "%v(%v(v))" $to $proxyFunc }}
}
{{ end }}

{{ define "converterFunc" }}
{{- $to := index . 0 -}}
{{- $parseFuncs := index . 1 -}}
{{- $floatCheck := index . 2 -}}
{{- $thisFunc := Title $to -}}
// {{$thisFunc}} converts the given value to a {{$to}}.
func {{$thisFunc}}(v interface{}) {{$to}} {
	v = Indirect(v)

	switch T := v.(type) {
	case {{$thisFunc}}Converter:
		if T != nil {
			return T.{{$thisFunc}}(T)
		}
	case bool:
		if T != false {
			return 1
		}
	case string:
		{{- range $parseFunc := $parseFuncs }}
			if parsed, err := strconv.{{ $parseFunc }}; err == nil {
			{{- if eq $parseFunc "ParseFloat(T, 64)" }}
			{{- if eq $floatCheck "true" }}
				if 1 > parsed {
					return {{ $to }}(0)
				}
			{{- end }}
			{{- end }}
			return {{ $to }}(parsed)
			}
		{{- end }}
		if parsed, err := strconv.ParseBool(T); err == nil {
			if parsed {
				return {{ $to }}(1)
			}
			return 0
		}
	case time.Duration:
		return {{ $to }}(T)
	case time.Time:
		return {{ $to }}(T.Unix())
	{{- range $from := (args "complex64" "complex128") }}
	case {{ $from }}:
		return {{ Convert $from $to "real(T)" }}
	{{- end }}
	{{- range $from := (args "uint" "uint8" "uint16" "uint32" "uint64" "float32") }}
	case {{ $from }}:
		return {{ Convert $from $to "T" }}
	{{- end }}
	{{- range $from := (args "float64" "int" "int8" "int16" "int32" "int64") }}
	case {{ $from }}:
		return {{ Convert $from $to "T" }}
	{{- end }}
	}
	return 0
}
{{ end }}

{{ define "converterInterface" }}
{{- $to := index . 0 -}}
{{- $func := index . 0 | Title }}
// {{ $func }}Converter interface allows a value to be converted to a {{$to}}.
type {{ $func }}Converter interface {
	{{ $func }}(from interface{}) {{ $to }}
}
{{- end -}}
