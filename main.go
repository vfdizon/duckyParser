package main

import (
	"bufio"
	"flag"
	"os"
)

var inputFileFlag string
var outputFileFlag string

func init() {
	flag.StringVar(&inputFileFlag, "f", "", "The file path of the file containing the desired script")
	flag.StringVar(&outputFileFlag, "o", "", "The file path of the desired output file (with no extension)")
}

func main() {
	flag.Parse()
	lines := parseFile()
	createFile(outputFileFlag+".txt", lines)

}

func createFile(name string, input []string) {
	f, _ := os.Create(name)
	defer f.Close()

	for _, val := range input {
		f.Write([]byte(val + "\n"))
	}
}

func parseFile() []string {
	ret_val := []string{}
	file, err := os.Open(inputFileFlag)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currentLine := scanner.Text()
		if currentLine != "" {
			ret_val = append(ret_val, "STRING "+currentLine)
			ret_val = append(ret_val, "ENTER")
		} else {
			ret_val = append(ret_val, "ENTER")
		}
	}
	return ret_val
}
