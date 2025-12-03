package main

import "fmt"

func main(){

	pods := []string{"pod-a","pod-b","pod-c"}
	
	for _, p := range pods{
		fmt.Println("Pod", p)
	}

	nodeStatus := map[string]string{
		"node1": "Ready",
		"node2": "NotReady",
	}	

	for key, value := range nodeStatus{
		if value == "NotReady" {

			fmt.Println("Node: ", key)
		}
	}

	return 
}