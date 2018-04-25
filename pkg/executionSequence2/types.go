package executionSequence2

import (
	"fmt"

	"github.com/yangyumo123/k8s-demo/pkg/executionSequence"
	"github.com/yangyumo123/k8s-demo/pkg/executionSequence3"
)

// EsConst test
const EsConst = "esConst"

// EsVar test
var EsVar = "esVar"

func init() {
	fmt.Println("executionSequence2-types.go-init " + EsConst)
	fmt.Println("executionSequence-types.go-init " + EsVar)
	executionSequence3.Demo()
	executionSequence.Demo()
}
