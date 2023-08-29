package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func main() {
	fmt.Println("[servers]")
	records := readCsvFile("./input/1.csv")
	// test := records[1][8]
	// println(test)
	count := 0

	for _, line := range records {
		// a := records[len(records)-1]
		// fmt.Println("debug hui", line)
		if count == 0 {
			count++
			continue

		}
		count++
		if line[1] == "cdp" {
			if len(line[7]) > 0 {
				// fmt.Println("debug len ", len(line))
				// fmt.Println(strings.TrimSuffix(line[0], "\n"), "-", strings.TrimSuffix(line[8], "\n"))
				// fmt.Println("ansible_host=", line[8])
				tempString := strings.TrimSuffix(line[0], "\n") + "-" + strings.TrimSuffix(line[7], "\n") + "-" + strings.TrimSuffix(line[8], "\n")
				tempString1 := strings.TrimSpace(tempString)
				tempString2 := tempString1 + " ansible_host=" + line[7]
				fmt.Println(tempString2)
			} else {

				// fmt.Println(strings.TrimSuffix(line[0], "\n"), "-", strings.TrimSuffix(line[7], "\n"))
				// fmt.Println("ansible_host=", line[7])
				tempString := strings.TrimSuffix(line[0], "\n") + "-" + strings.TrimSuffix(line[8], "\n")
				tempString1 := strings.TrimSpace(tempString)
				tempString2 := tempString1 + " ansible_host=" + line[8]
				fmt.Println(tempString2)
			}
		}
	}
	// temp = template.Must(template.ParseFiles("hosts.tpl"))
	// file, _ := os.Create("output/hosts")
	// defer file.Close()

	// // apply the template to the vars map and write the result to file.
	// err := temp.Execute(file,)
	// err := temp.Execute(file, TempServiceList)

	// fmt.Println("hello")
}
