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

func GetLogLine(lineLine string, value string, text string) LogLine {
	logLine := LogLine{lineLine, value, text}
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

type LogLine struct {
	typeLine, valueLine, textLine string
}
