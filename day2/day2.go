package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile() [][]int {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines [][]int

	for fileScanner.Scan() {
		line := fileScanner.Text()
		var arr []int
		for _, e := range strings.Split(line, " ") {
			value, err := strconv.Atoi(e)
			if err != nil {
				log.Fatal(err)
			}
			arr = append(arr, value)
		}
		fileLines = append(fileLines, arr)
	}
	file.Close()

	return fileLines
}
func main() {
	values := readFile()
	safeReports := 0
	// part one
	for _, report := range values {
		if report[0] != report[1] {
			isIncreasing := report[0] < report[1]
			isSafe := true
			for index, level := range report {
				if index != 0 {
					previousLevel := report[index-1]
					switch {
					case isIncreasing && level > previousLevel && level-previousLevel < 4:
						{
							// ok
						}
					case !isIncreasing && previousLevel > level && previousLevel-level < 4:
						{
							// ok
						}
					default:
						isSafe = false
					}
				}
			}
			if isSafe {
				safeReports++
			}
		}
	}
	fmt.Println("PART ONE =>", safeReports)

	//part two
	reallySafe := 0
	for _, reportTmp := range values {
		reportLength := len(reportTmp)
		isFullySafe := false
		i := 0
		for i < reportLength {
			var report []int
			report = append([]int{}, reportTmp[:i]...)  // copie la partie avant i
			report = append(report, reportTmp[i+1:]...) // ajoute la partie après i
			i++
			if report[0] != report[1] {
				isIncreasing := report[0] < report[1]
				isSafe := true
				for index, level := range report {
					if index != 0 {
						previousLevel := report[index-1]
						switch {
						case isIncreasing && level > previousLevel && level-previousLevel < 4:
							{
								// ok
							}
						case !isIncreasing && previousLevel > level && previousLevel-level < 4:
							{
								// ok
							}
						default:
							isSafe = false
						}
					}
				}
				if isSafe {
					isFullySafe = true
				}
			}
		}
		if isFullySafe {
			reallySafe++
		}
	}
	fmt.Println("PART TWO =>", reallySafe)
}
