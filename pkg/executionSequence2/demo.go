package executionSequence2

import (
	"fmt"

	"github.com/yangyumo123/k8s-demo/pkg/executionSequence3"
)

func init() {
	fmt.Println("executionSequence2-demo.go-init demo")
	executionSequence3.Demo()
}

// Demo test 2
func Demo() {
	fmt.Println("executionSequence2-demo.go Demo()")
}
