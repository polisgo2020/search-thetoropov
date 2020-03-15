package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main()  {
	var args []string = os.Args //outputFileName textFileName textFileName ...
	if len(args) < 3 {
		log.Fatal("Not enough input data")
	}
	text := fileToText(args)
	index := index(text)
	invIndex := invertedIndex(index)
	writeToFile(args, invIndex)
}

func fileToText(args []string) []string{
	text := make([]string, len(args) - 2)
	for i := 0; i < len(args) - 2; i++ {
		data, err := ioutil.ReadFile(args[i + 2])
		checkError(err)
		text[i] = string(data)
	}
	return text
}

func index(text []string) [][]string {
	words := make([][]string, len(text))
	for i := 0; i < len(text); i++ {
		words[i] = strings.Fields(text[i])
	}
	return words
}

func invertedIndex(words [][]string) map[string]map[int]int {
	invInd := make(map[string]map[int]int)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if invInd[words[i][j]] != nil {
				index := invInd[words[i][j]]
				index[i] = i
			} else {
				invInd[words[i][j]] = make(map[int]int)
				invInd[words[i][j]][i] = i
			}
		}
	}
	return invInd
}

func writeToFile(args []string, invInd map[string]map[int]int) {
	file, err := os.Create(args[1])
	checkError(err)
	defer file.Close()
	for key, value := range invInd {
		var tmp string
		for k:= range value {
			tmp += strconv.Itoa(k) + " "
		}
		_, err := file.WriteString(key + " " + tmp + "\n")
		checkError(err)
	}
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}