package runtime

type ResourceVersioner interface {
	ResourceVersion(obj Object) (uint64, error)
	SetResourceVersion(obj Object, version uint64) error
}
type Object interface {
	IsAnAPIObject()
}
type Codec interface {
	Encode(obj Object) ([]byte, error)
	Decode(data []byte) (Object, error)
	DecodeInto(data []byte, obj Object) error
}
