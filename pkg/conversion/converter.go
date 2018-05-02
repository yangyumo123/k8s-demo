package conversion

import (
	"reflect"
)

type typePair struct {
	source reflect.Type
	dest   reflect.Type
}
type Converter struct {
	funcs map[typePair]reflect.Value
	Debug DebugLogger
}
type DebugLogger interface {
	Logf(format string, args ...interface{})
}

func NewConverter() *Converter {
	return &Converter{
		funcs: map[typePair]reflect.Value{},
	}
}
