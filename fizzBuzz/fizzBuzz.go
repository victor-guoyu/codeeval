package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type config struct {
	X, Y, N int
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	return string(data)
}

func parseContent(content string) []config {
	var configs []config
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 3 {
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			n, _ := strconv.Atoi(parts[2])
			configs = append(configs, config{x, y, n})
		}
	}
	return configs
}

func fizzBuzz(config config) {
	result := make([]string, config.N)
	for i := 1; i <= config.N; i++ {
		if i%config.X == 0 {
			result[i-1] += "F"
		}
		if i%config.Y == 0 {
			result[i-1] += "B"
		}
		if len(result[i-1]) == 0 {
			result[i-1] = strconv.Itoa(i)
		}
	}
	fmt.Println(strings.Join(result, " "))
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Println("Please specify the input file")
		return
	}
	content := readFile(flag.Args()[0])
	configs := parseContent(content)
	for _, config := range configs {
		fizzBuzz(config)
	}
}
