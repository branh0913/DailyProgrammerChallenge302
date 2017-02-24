package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"strings"
	"bufio"
)

func GetAtomicSymLen(fileName string, symLen int) []map[string]string {
	elementFile, err := os.Open(fileName)

	if err != nil{
		panic("file cannot be opened!")
	}

	atomicSymbol := make([]map[string]string, 0)

	csvFile := csv.NewReader(elementFile)

	fileContent,_ := csvFile.ReadAll()

	for i := range fileContent{
		if len(fileContent[i][1]) == symLen{

			atomicSymbol = append(atomicSymbol,
				map[string]string{strings.ToLower(
					fileContent[i][1]): strings.ToLower(fileContent[i][0])})

		}
	}

	return atomicSymbol


}

func MatchDouble(str string, doubleEvalMap []map[string]string) string{

	for i := range doubleEvalMap {
		if val, ok := doubleEvalMap[i][str]; ok{
			return val
		}

	}
	return "fail"
}

func MatchSingle(singeStr string, singleEvalMap []map[string]string ) string {

	var result string
	for i := range singleEvalMap{
		if val, ok := singleEvalMap[i][singeStr]; ok{
			result = val
		}
	}
	return result


}

func main() {

	inputString := "C:\\Users\\BHARRELL\\GolangProjects\\src\\github.com\\branh0913\\ProgrammerChallenge305\\input.txt"
	fileString := "C:\\Users\\BHARRELL\\GolangProjects\\src\\github.com\\branh0913\\ProgrammerChallenge305\\element.csv"



	doubleList := GetAtomicSymLen(fileString, 2)
	singleList := GetAtomicSymLen(fileString, 1)



	inputFile, err := os.Open(inputString)

	if err != nil{
		panic("Cannot open file")
	}
	defer inputFile.Close()
	parseMap := make(map[string][]string, 0)

	inputScanner := bufio.NewScanner(inputFile)

	for inputScanner.Scan() {
		fileLine := strings.TrimSpace(string(inputScanner.Bytes()))
		resultList := make([]string,0)
		filterList := make([]string, 0)




		for i := 0; i < len(fileLine) - 1; i = i + 2 {
			charPair := string(fileLine[i]) + string(fileLine[i + 1])


			doubleCheck := MatchDouble(charPair, doubleList)
			resultList = append(resultList, doubleCheck)

			if doubleCheck == "fail" {
				for ch := range charPair {
					singleCheck := MatchSingle(string(charPair[ch]), singleList)
					if singleCheck == "fail"{
						fmt.Println("fail won't be added")
					} else{
						resultList = append(resultList, singleCheck)
					}

				}
			}


		}
		for results := range resultList{
			if resultList[results] != "fail"{
				filterList = append(filterList, resultList[results])

			}
		}
		parseMap[fileLine] = filterList

	}
	for k,v := range parseMap{
		fmt.Println(k, "("+strings.Join(v, ", ")+")")
	}


}
