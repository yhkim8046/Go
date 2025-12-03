package main

import "fmt"
import "os/exec"

func main(){
	fmt.Println("first go programme")

	cmd := exec.Command("bash","-c","kubectl get nodes -o wide")
	output, err := cmd.CombinedOutput()
	if err != nil{
		fmt.Println("Error", err)
		return
	}
	fmt.Println(string(output))
}