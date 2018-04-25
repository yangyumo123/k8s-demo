package main

import (
	"fmt"

	"github.com/yangyumo123/k8s-demo/pkg/executionSequence"
	_ "github.com/yangyumo123/k8s-demo/pkg/executionSequence2"
)

const mainConstA = "ca"
const mainConstB = "cb"

func init() {
	var mainInitVar = mainInitVarFunc("demo")
	fmt.Println(mainInitVar)
	const mainInitConst string = "demo"
	fmt.Println(mainInitConst)
	fmt.Println(mainVarA)
}

var mainVarA = "va"
var mainVarB = "vb"

func mainInitVarFunc(demo string) string {
	return demo
}

func main() {
	fmt.Println(mainVarA)
	fmt.Println(mainVarB)
	fmt.Println(mainConstA)
	fmt.Println(mainConstB)
	executionSequence.Demo()
	// fmt.Println(executionSequence2.EsConst)
	// executionSequence2.Demo()
}
