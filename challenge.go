package main

import "fmt"
import "os/exec"
import "strconv"
import "strings"

func podCounter(threshold int)bool{
	cmd := exec.Command("bash","-c","kubectl get pods -A --no-headers | wc -l")

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	countStr := strings.TrimSpace(string(out))
	podCount, err := strconv.Atoi(countStr)
	if err != nil {
		fmt.Println("failed: ", err)
		return false
	}

	return podCount >= threshold
}


func main(){
	result := podCounter(15)
	fmt.Println(result)
}