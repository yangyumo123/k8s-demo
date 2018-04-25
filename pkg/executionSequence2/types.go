package executionSequence2

import (
	"fmt"

	"github.com/yangyumo123/k8s-demo/pkg/executionSequence"
)

const EsConst = "esConst"

var EsVar = "esVar"

func init() {
	fmt.Println(EsConst)
	fmt.Println(EsVar)
	executionSequence.Demo()
}
