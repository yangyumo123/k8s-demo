package executionSequence

import (
	"fmt"
)

const esConst = "esConst"

var esVar = "esVar"

func init() {
	fmt.Println("executionSequence-types.go-init " + esConst)
}
func init() {
	fmt.Println("executionSequence-types.go-init " + esVar)
}
