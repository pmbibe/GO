package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Req struct {
	question string
	answer   string
}

func main() {
	var Reqq []Req
	csvFile, _ := os.Open("problems.csv")
	defer csvFile.Close()
	csvLines, _ := csv.NewReader(csvFile).ReadAll()
	for _, line := range csvLines {
		X := Req{
			question: line[0],
			answer:   line[1],
		}
		Reqq = append(Reqq, X)
	}
	i := 0
	var ansList []string
	var ans string
	for i < len(Reqq) {
		fmt.Printf("What is answer of question: %v ? \n", Reqq[i].question)
		fmt.Print("Answer:")
		fmt.Scan(&ans)
		ansList = append(ansList, ans)
		i = i + 1
	}
	point := 0
	j := 0
	for j < len(Reqq) {
		if ansList[j] == Reqq[j].answer {
			point = point + 1
		}
		j = j + 1
	}

	fmt.Printf("Your point: %d", point)

}
