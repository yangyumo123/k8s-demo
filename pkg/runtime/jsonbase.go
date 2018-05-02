package runtime

func NewJSONBaseResourceVersioner() ResourceVersioner {
	return &jsonBaseResourceVersioner{}
}

type jsonBaseResourceVersioner struct{}

func (jsonBaseResourceVersioner) ResourceVersion(obj Object) (uint64, error) {
	json, err := FindJSONBase(obj)
	if err != nil {
		return 0, err
	}
	return json.ResourceVersion(), nil
}
func (jsonBaseResourceVersioner) SetResourceVersion(obj Object, version uint64) error {
	json, err := FindJSONBase(obj)
	if err != nil {
		return err
	}
	json.SetResourceVersion(version)
	return nil
}

type JSONBaseInterface interface {
	ID() string
	SetID(ID string)
	Kind() string
	SetKind(kind string)
	ResourceVersion() uint64
	SetResourceVersion(version uint64)
	APIVersion() string
	SetAPIVersion(version string)
}
type genericJSONBase struct {
	id              *string
	kind            *string
	apiVersion      *string
	resourceVersion *uint64
}

func (g genericJSONBase) ID() string {
	return *g.id
}
func (g genericJSONBase) SetID(ID string) {
	*g.id = ID
}
func (g genericJSONBase) Kind() string {
	return *g.kind
}
func (g genericJSONBase) SetKind(kind string) {
	*g.kind = kind
}
func (g genericJSONBase) APIVersion() string {
	return *g.apiVersion
}
func (g genericJSONBase) SetAPIVersion(version string) {
	*g.apiVersion = version
}
func (g genericJSONBase) ResourceVersion() uint64 {
	return *g.resourceVersion
}
func (g genericJSONBase) SetResourceVersion(version uint64) {
	*g.resourceVersion = version
}
