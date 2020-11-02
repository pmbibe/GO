package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//ReadingFileIntoMemory Reading an entire file into memory
func ReadingFileIntoMemory(file string) []string {
	var fileDic []string
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	dataString := strings.Split(string(data), "\n")
	for _, s := range dataString {
		fileDic = append(fileDic, s)

	}
	return fileDic
}

//ReadingFileInBytes o o
func ReadingFileInBytes(file string) {
	data, err := os.Open(file)
	defer data.Close()
	if err != nil {
		fmt.Println(err)
	}
	r := bufio.NewReader(data)
	b := make([]byte, 100)
	for {
		n, err := r.Read(b)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(b[0:n]))
	}
}

//ReadingFileLineByLine Reading a file line by line
func ReadingFileLineByLine(file string) []string {
	var fileDic []string
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		fileDic = append(fileDic, s.Text())
	}
	err = s.Err()
	if err != nil {
		fmt.Println(err)
	}
	return fileDic

}

//Convert all to Lower case
func Convert(s string) string {
	return strings.ToLower(s)
}

//AnalystString Find and count the number of characterer in string
func AnalystString(s1 string) map[string]int {
	s1 = Convert(s1)
	duplicateFrequency := make(map[string]int)
	for _, item := range s1 {
		item := string(item)
		_, exist := duplicateFrequency[item] //Check if exist or not
		if exist {
			duplicateFrequency[item]++ // increase counter by 1 if already in the map
		} else {
			duplicateFrequency[item] = 1 // else start counting from 1
		}
	}
	return duplicateFrequency
}

//CompareString o o
func CompareString(s1, s2 string) bool {
	s1 = Convert(s1)
	s2 = Convert(s2)
	s1S := AnalystString(s1)
	s2S := AnalystString(s2)
	if len(s1S) < len(s2S) {
		for i := range s1S {
			if s1S[i] != s2S[i] {
				return false
			}
		}
	} else {
		for i := range s2S {
			if s1S[i] != s2S[i] {
				return false
			}
		}
	}

	return true
}

//FindAnagrams o o
func FindAnagrams(dictionary []string, word string) []string {
	var stringArr []string
	word = Convert(word)
	for _, value := range dictionary {
		if CompareString(value, word) {
			stringArr = append(stringArr, value)
		}
	}
	fmt.Println(stringArr)
	return stringArr
}

func main() {
	FindAnagrams(ReadingFileLineByLine("dictionary.txt"), "The eyes")
	// fmt.Println(Convert("Dictionary"))
}
