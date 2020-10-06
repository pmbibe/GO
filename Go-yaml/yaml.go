package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

type Node struct {
	name  string
	fdqn  string
	ipv4  string
	ip_v4 string
}
type Nodes struct {
	name string
}

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
	var b Nodes
	err := yaml.Unmarshal([]byte(data), &b)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	fmt.Printf("%#v\n", b)

}
