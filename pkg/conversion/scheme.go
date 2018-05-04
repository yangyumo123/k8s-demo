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

// AddKnownTypes register interface{}
func (s *Scheme) AddKnownTypes(version string, types ...interface{}) {
	knownTypes, found := s.versionMap[version]
	if !found {
		knownTypes = map[string]reflect.Type{}
		s.versionMap[version] = knownTypes
	}
	for _, value := range types {
		t := reflect.TypeOf(value)
		if t.Kind() != reflect.Ptr {
			panic("")
		}
		t = t.Elem()
		if t.Kind() != reflect.Struct {
			panic("")
		}
		knownTypes[t.Name()] = t
		s.typeToVersion[t] = version
	}
}
