package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type Req struct {
	question string
	answer   string
}

func readFile(file string) []Req {
	var Reqq []Req
	csvFile, _ := os.Open(file)
	defer csvFile.Close()
	csvLines, _ := csv.NewReader(csvFile).ReadAll()
	for _, line := range csvLines {
		X := Req{
			question: line[0],
			answer:   line[1],
		}
		Reqq = append(Reqq, X)
	}
	return Reqq
}
func getAnswer(file string, timeLimit int) []string {
	i := 0
	var ansList []string
	var ans string
	go func() {
		for i < len(readFile(file)) {
			fmt.Printf("What is answer of question: %v ? \n", readFile(file)[i].question)
			fmt.Print("Answer: ")
			fmt.Scan(&ans)
			ansList = append(ansList, ans)
			i = i + 1
		}
	}()
	time.Sleep(time.Duration(timeLimit) * time.Second)

	return ansList

}

func getPoint(file string, timeLimit int) {
	point := 0
	listAnswer := getAnswer(file, timeLimit)
	j := 0
	for j < len(listAnswer) {
		if listAnswer[j] == readFile(file)[j].answer {
			point = point + 1
		}
		j = j + 1
	}
	fmt.Print("\nYour score is: ", point)
}

func channelA(point int) {
	channelAws := make(chan int)
	defer close(channelAws)
}

func main() {
	fileInput := flag.String("csv", "problems.csv", "This is csv file")
	timeLimit := flag.Int("time", 5, "This is time limit")
	flag.Parse()
	file := fileInput
	getPoint(*file, *timeLimit)
}
