package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func detectNotReadyNode() [][2]string {
	cmd := exec.Command("bash", "-c", "kubectl get nodes --no-headers")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")

	var nodeList [][2]string

	for _, line := range lines { 
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		name := fields[0]
		status := fields[1]
		nodeList = append(nodeList, [2]string{name, status})
	}

	return nodeList
}

func main() {
	nodes := detectNotReadyNode()
	for _, node := range nodes {
		fmt.Printf("노드: %-10s 상태: %s\n", node[0], node[1])
	}
}
