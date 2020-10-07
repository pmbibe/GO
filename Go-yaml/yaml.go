package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

//NodeLevel0 OK
type NodeLevel0 struct {
	Node map[string]string
}

//NodeLevel OK
type NodeLevel struct {
	Node map[string]map[string]string `yaml:"node"`
}

// type NodeLevel1 struct {
// 	Node map[string]string
// }

var data = `
node:
    nodeA:
        fqdn: "TestScript"
        ipv4: "192.168.141.240/24"
        ip_v4: "192.168.141.240"
    nodeB:
        fqdn: "ComputeNode"
        ipv4: "192.168.141.242/24"
        ip_v4: "192.168.141.242"  
    nodeC:
        fqdn: "ControllerNode"
        ipv4: "192.168.141.243/24"
        ip_v4: "192.168.141.243" 
`

func main() {
	var b NodeLevel
	// var d []byte
	err := yaml.Unmarshal([]byte(data), &b)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	keys := make([]string, 0, len(b.Node))
	values := make([]map[string]string, 0, len(b.Node))

	for k, v := range b.Node {
		keys = append(keys, k)
		values = append(values, v)
		fmt.Print(k)
		fmt.Print("\n")
		fmt.Print(b.Node[k])
		fmt.Print("\n")
	}
	// fmt.Print(keys)
	// fmt.Print(values)
	// for i in keys {
	// 	fmt.Print(b.Node["nodeA"])
	// }

}
