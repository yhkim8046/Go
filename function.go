package main

import (
	"fmt"
	"strings"
)

func isReady(status string) bool {
	return status == "Ready"
}

func addAndMul(a, b int) (int, int){
	return a+b, a*b
}

func splitNameVersion(s string) (name, version string){
	parts := strings.SplitN(s, ":", 2)	
	if len(parts) == 2{
		name, version = parts[0], parts[1]
	} else {
		name = s 
		version = "unknown"
	}

	return
}

func main() {
	fmt.Println(isReady("Ready"))
	fmt.Println(isReady("False"))

	a := 5 
	b := 7

	fmt.Println(addAndMul(a, b))
	fmt.Println(splitNameVersion("kubelet: v1.29.3"))
}
