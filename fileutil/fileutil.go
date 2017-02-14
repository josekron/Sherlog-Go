package fileutil

//@Author: josekron
//Manage files

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func IsLocalFile(localFile string) bool {
	var isOk bool = true
	file, err := os.Open(localFile)
	if err != nil {
		fmt.Println("Error! The system cannot find the file: ", localFile)
		isOk = false
	}
	defer file.Close()
	return isOk
}

func IsValidFileExtension(localFile string) bool {
	var isValid bool = false
	var extensionFile = GetFileExtension(localFile)
	switch extensionFile {
	// case ".txt":
	// 	isValid = true
	case ".log":
		isValid = true
	default:
		isValid = false
		fmt.Println("Error! The extension  ", extensionFile, " is not valid")
	}
	return isValid
}

func GetFileExtension(localFile string) string {
	var validID = regexp.MustCompile(`\.([0-9a-z]+)(?:[\?#]|$)`)
	var extensionFile string = ".invalid"
	if validID.MatchString(localFile) {
		extensionFile = validID.FindString(localFile)
	}
	return extensionFile
}

func SearchInFile(localFile string, text string) []LogLine {
	file, err := os.Open(localFile)
	if err != nil {
		fmt.Println("Error! The system cannot find the file: ", localFile)
	}
	defer file.Close()

	//Map:
	logList := []LogLine{}
	rpDate := regexp.MustCompile("\\d{4,4}-\\d{2,2}-\\d{2,2} \\d{1,2}:\\d{2,2}:\\d{2,2}")

	text = strings.ToUpper(strings.TrimSpace(text))
	reader := bufio.NewReader(file)
	var line string
	var counterLine int = 1
	for {
		line, err = reader.ReadString('\n')

		var lineLog = strings.ToUpper(line)
		if strings.Contains(lineLog, text) {

			//regex for search date in lineLog
			ocurrences := rpDate.FindString(lineLog)

			logLine := GetLogLine("line", strconv.Itoa(counterLine), line)
			if len(ocurrences) > 0 {
				logLine = GetLogLine("date", string(ocurrences), line)
			}
			logList = AddLogLine(logList, logLine)
		}

		counterLine++

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}
	return logList
}

func GetLogLine(typeLine string, value string, text string) LogLine {
	logLine := LogLine{typeLine, value, text}
	return logLine
}

func AddLogLine(lines []LogLine, line LogLine) []LogLine {
	newLines := []LogLine{line}
	for _, f := range newLines {
		lines = append(lines, f)
	}
	return lines
}

func GetLineType(logLine *LogLine) string {
	return logLine.typeLine
}

func GetLineValue(logLine *LogLine) string {
	return logLine.valueLine
}

func GetText(logLine *LogLine) string {
	return logLine.textLine
}

func PrintLogLine(logLine *LogLine) {
	if logLine.typeLine == "date" {
		fmt.Println(logLine.valueLine, " -> ", logLine.textLine)
	} else {
		fmt.Println(logLine.typeLine+" "+logLine.valueLine, " -> ", logLine.textLine)
	}
}

func PrintLogLineList(logsLineList [][]LogLine, fileLogs []string) {

	counter := make([]int, len(logsLineList))
	for i := 0; i < len(counter); i++ {
		counter[i] = len(logsLineList[i]) - 1
	}

	var proccesed bool = false
	prevSelectedLogLine := GetLogLine("line", "-1", "")

	for !proccesed {
		selectedLogLine := GetLogLine("line", "-1", "")
		selectedLog := 0

		for i := 0; i < len(logsLineList); i++ {

			if counter[i] > 0 {

				if selectedLogLine.valueLine != "-1" {

					//Check if previous line was a type=date and if current line has the same date
					if logsLineList[i][len(logsLineList[i])-counter[i]].typeLine == "date" && prevSelectedLogLine.valueLine != "-1" && prevSelectedLogLine.typeLine == "date" {

						if logsLineList[i][len(logsLineList[i])-counter[i]].valueLine == prevSelectedLogLine.valueLine {
							selectedLogLine = logsLineList[i][len(logsLineList[i])-counter[i]]
							selectedLog = i
							i = len(logsLineList)
						}

					} else {

						//Preference type=line
						if logsLineList[i][len(logsLineList[i])-counter[i]].typeLine == "line" {
							selectedLogLine = logsLineList[i][len(logsLineList[i])-counter[i]]
							selectedLog = i
						} else if selectedLogLine.typeLine != "line" {
							selectedLogLine = logsLineList[i][len(logsLineList[i])-counter[i]]
							selectedLog = i
						}
					}
				} else { //First line
					selectedLogLine = logsLineList[i][len(logsLineList[i])-counter[i]]
					selectedLog = i
				}
			}
		}

		fmt.Print("[", fileLogs[selectedLog], "] - ")
		PrintLogLine(&logsLineList[selectedLog][len(logsLineList[selectedLog])-counter[selectedLog]])

		prevSelectedLogLine = logsLineList[selectedLog][len(logsLineList[selectedLog])-counter[selectedLog]]
		selectedLogLine = GetLogLine("line", "-1", "")
		counter[selectedLog]--

		//check if all lines are processed:
		proccesed = true
		for i := 0; i < len(counter); i++ {
			if counter[i] > 0 {
				proccesed = false
				break
			}
		}
	}
}

type LogLine struct {
	typeLine, valueLine, textLine string
}
