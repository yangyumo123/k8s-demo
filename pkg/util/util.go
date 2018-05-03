package util

var ReallyCrash bool

type IntstrKind int

const (
	IntstrInt IntstrKind = iota
	IntstrString
)

type IntOrString struct {
	Kind   IntstrKind
	IntVal int
	StrVal string
}
