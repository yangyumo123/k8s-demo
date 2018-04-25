package main

import (
	"fmt"

	// "github.com/yangyumo123/k8s-demo/pkg/executionSequence"
	_ "github.com/yangyumo123/k8s-demo/pkg/executionSequence2"
)

var mainVar = "v"

const mainConst = "c"

func init() {
	var mainInitVar = mainInitVarFunc("v")
	const mainInitConst = "c"
	fmt.Println("main-init " + mainInitVar)
	fmt.Println("main-init " + mainInitConst)
}

func mainInitVarFunc(demo string) string {
	return demo
}

func main() {
	fmt.Println("main " + mainVar)
	fmt.Println("main " + mainConst)
	// executionSequence.Demo()
	// fmt.Println(executionSequence2.EsConst)      //undefined: executionSequence2
	// executionSequence2.Demo()                    //undefined: executionSequence2
}
