package runtime

import (
	"fmt"
	"reflect"

	"github.com/yangyumo123/k8s-demo/pkg/conversion"
)

var DefaultResourceVersioner ResourceVersioner = NewJSONBaseResourceVersioner()
var DefaultScheme = NewScheme("", "v1beta1")
var DefaultCodec Codec = DefaultScheme

type Scheme struct {
	raw *conversion.Scheme
}

func NewScheme(internalVersion string, externalVersion string) *Scheme {
	s := &Scheme{conversion.NewScheme()}
	s.raw.InternalVersion = internalVersion
	s.raw.ExternalVersion = externalVersion
	s.raw.MetaInsertionFactory = metaInsertion{}
}

type metaInsertion struct {
	JSONBase struct {
		APIVersion string `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
		Kind       string `json:"kind,omitempty" yaml:"kind,omitempty"`
	} `json:",inline" yaml:",inline"`
}

func (metaInsertion) Create(version, kind string) interface{} {
	m := metaInsertion{}
	m.JSONBase.APIVersion = version
	m.JSONBase.Kind = kind
	return &m
}
func (metaInsertion) Interpret(in interface{}) (version, kind string) {
	m, ok := in.(*metaInsertion)
	if !ok {
		return "", ""
	}
	return m.JSONBase.APIVersion, m.JSONBase.Kind
}

// FindJSONBase test
func FindJSONBase(obj Object) (JSONBaseInterface, error) {
	v, err := enforcePtr(obj)
	if err != nil {
		return nil, fmt.Errorf("couldn't get ptr")
	}
	t := v.Type()
	name := t.Name()
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, but got %v: %v (%#v)", v.Kind(), name, v.Interface())
	}
	jsonBase := v.FieldByName("JSONBase")
	if !jsonBase.IsValid() {
		return nil, err
	}
	g, err := newGenericJSONBase(jsonBase)
	if err != nil {
		return nil, err
	}
	return g, nil
}
func enforcePtr(obj Object) (reflect.Value, error) {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr {
		return reflect.Value{}, fmt.Errorf("")
	}
	return v.Elem(), nil
}
func newGenericJSONBase(v reflect.Value) (genericJSONBase, error) {
	g := genericJSONBase{}
	if err := fieldPtr(v, "ID", &g.id); err != nil {
		return g, err
	}
	if err := fieldPtr(v, "Kind", &g.kind); err != nil {
		return g, err
	}
	if err := fieldPtr(v, "APIVersion", &g.apiVersion); err != nil {
		return g, err
	}
	if err := fieldPtr(v, "ResourceVersion", &g.resourceVersion); err != nil {
		return g, err
	}
	return g, nil
}
func fieldPtr(v reflect.Value, fieldName string, dest interface{}) error {
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("")
	}
	v = reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("")
	}
	v = v.Elem()
	field = field.Addr()
	if field.Type().AssignableTo(v.Type()) {
		v.Set(field)
		return nil
	}
	if field.Type().ConvertibleTo(v.Type()) {
		v.Set(field.Convert(v.Type()))
		return nil
	}
	return fmt.Errorf("")
}

// AddKnownTypes register Object
func (s *Scheme) AddKnownTypes(version string, types ...Object) {
	interfaces := make([]interface{}, len(types))
	for i := range interfaces {
		interfaces[i] = types[i]
	}
	s.raw.AddKnownTypes(version, interfaces)
}
