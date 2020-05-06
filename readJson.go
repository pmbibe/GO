package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

type Server struct {
	NAMESERVER string
	SERVICE    []string
}

var lastStatusService uint

func checkServiceRunning(service string, server string) {
	serviceName := "./exitCode.sh " + service + " " + server + " ;echo $?"
	StatusCode := exec.Command("sh", "-c", serviceName)
	statusCode, _ := StatusCode.Output()
	sttCode := string(statusCode)
	if sttCode == "0\n" && (lastStatusService != 0) {
		log.Print("Service is running")
		lastStatusService = 0
	} else if sttCode != "0\n" && (lastStatusService != 1) {
		log.Print("Service is Dead")
		lastStatusService = 1
	}

}
func main() {
	data, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Print(err)
	}

	var obj []Server
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}
	for i := 0; i < len(obj); i++ {
		for j := 0; j < len(obj[i].SERVICE); j++ {
			go checkServiceRunning(obj[i].SERVICE[j], obj[i].NAMESERVER)
		}

	}

}
