package reflectDemo

// DemoJSONBase for test
type DemoJSONBase struct {
	ID              *string `json:"id"`
	Kind            *string `json:"kind"`
	APIVersion      *string `json:"apiVersion,omitempty"`
	ResourceVersion *uint64 `json:"resourceVersion,omitempty"`
	Env
}

// Env for test
type Env struct {
	Name  string
	Value string
}

// SetID set id value
func (d DemoJSONBase) SetID(id string) {
	*d.ID = id
}

// GetID get id value
func (d DemoJSONBase) GetID() string {
	return *d.ID
}
