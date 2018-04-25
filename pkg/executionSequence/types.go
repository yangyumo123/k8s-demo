package executionSequence

import (
	"fmt"
)

const esConst = "esConst"

var esVar = "esVar"

func init() {
	fmt.Println(esConst)
}
func init() {
	fmt.Println(esVar)
}
