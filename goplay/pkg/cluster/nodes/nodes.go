package nodes

import(
	"encoding/json"
	"os/exec"
	"strings"
)

type Node struct{
	Name string
	Status string
	IP string
}

// 공개 함수: kubectl -o json 파싱
func List()([]Node, error){
	cmd := exec.Commnad("bash","-c","kubectl get nodes -o json")

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var raw struct {
        Items []struct {
            Metadata struct {
                Name string `json:"name"`
            } `json:"metadata"`
            Status struct {
                Conditions []struct {
                    Type   string `json:"type"`
                    Status string `json:"status"`
                } `json:"conditions"`
                Addresses []struct {
                    Type    string `json:"type"`
                    Address string `json:"address"`
                } `json:"addresses"`
            } `json:"status"`
        } `json:"items"`
    }

	if err := json.Unmarshal(out, &raw); err != nil{
		return nil, err
	}

	var result []Node
    for _, it := range raw.Items {
        st := "Unknown"
        for _, c := range it.Status.Conditions {
            if c.Type == "Ready" {
                if strings.EqualFold(c.Status, "True") {
                    st = "Ready"
                } else {
                    st = "NotReady"
                }
                break
            }
        }
        ip := ""
        for _, a := range it.Status.Addresses {
            if a.Type == "InternalIP" {
                ip = a.Address
                break
            }
        }
        result = append(result, Node{Name: it.Metadata.Name, Status: st, IP: ip})
    }
    return result, nil
}

// 공개 함수: 필터링
func FilterNotReady(nodes []Node) []Node {
    var bad []Node
    for _, n := range nodes {
        if n.Status != "Ready" {
            bad = append(bad, n)
        }
    }
    return bad
}
