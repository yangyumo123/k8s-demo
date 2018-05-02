package conversion

import (
	"reflect"
)

type Scheme struct {
	versionMap           map[string]map[string]reflect.Type
	typeToVersion        map[reflect.Type]string
	converter            *Converter
	Indent               bool
	InternalVersion      string
	ExternalVersion      string
	MetaInsertionFactory MetaInsertionFactory
}

func NewScheme() *Scheme {
	return &Scheme{
		versionMap:           map[string]map[string]reflect.Type{},
		typeToVersion:        map[reflect.Type]string{},
		converter:            NewConverter(),
		InternalVersion:      "",
		ExternalVersion:      "v1",
		MetaInsertionFactory: metaInsertion{},
	}
}

type MetaInsertionFactory interface {
	Create(verison, kind string) interface{}
	Interpret(interface{}) (verison, kind string)
}

type metaInsertion struct {
	Version string `json:"verison,omitempty" yaml:"verison,omitempty"`
	Kind    string `json:"kind,omitempty" yaml:"kind,omitempty"`
}

func (metaInsertion) Create(verison, kind string) interface{} {
	m := metaInsertion{}
	m.Version = verison
	m.Kind = kind
	return &m
}
func (metaInsertion) Interpret(in interface{}) (version, kind string) {
	m := in.(*metaInsertion)
	return m.Version, m.Kind
}
