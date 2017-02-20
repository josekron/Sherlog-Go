package main

//@Author: josekron
//Main class

import (
	"bufio"
	"fileutil"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

func main() {
	title()

	//Variables:
	fileLogs := []string{}
	//fileLogs := []string{"C:\\server.log", "C:\\server2.log"}

	//Commands menu:
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for text != "exit" {
		fmt.Print("> ")
		scanner.Scan()

		text = scanner.Text()
		words := strings.Split(text, " ")

		switch words[0] {
		case "load":
			fileLogs = addFile(words, fileLogs)
			displayFileLogs(fileLogs)
		case "search":
			searchFiles(words, fileLogs)
		case "exit":
			fmt.Println("Bye!")
		default:
			fmt.Println("Command not recognized!")
		}
	}
}

func searchFiles(words []string, fileLogs []string) {
	if len(words) == 1 {
		fmt.Println("Write the text for the search")

	} else if len(fileLogs) == 0 {
		fmt.Println("Add at least one file!")

	} else {
		var text string
		for i := 1; i < len(words); i++ {
			text += words[i]
			if i < len(words) {
				text += " "
			}
		}

		//Create goroutines for each file:
		fileLogsList := make([][]fileutil.LogLine, len(fileLogs))
		var wg sync.WaitGroup

		for i := 0; i < len(fileLogs); i++ {
			fmt.Println("Starting search in ", fileLogs[i])
			wg.Add(1)
			go workerSearchInFile(&wg, fileLogsList, i, fileLogs[i], text)
		}
		fmt.Println("Waiting for the searches to finish")
		wg.Wait()
		fmt.Println("Search done!")

		fileutil.PrintLogLineList(fileLogsList, fileLogs)
	}
}

func workerSearchInFile(wg *sync.WaitGroup, fileLogsList [][]fileutil.LogLine, pos int, fileLog string, text string) {
	defer wg.Done()
	logList := fileutil.SearchInFile(fileLog, text)
	fileLogsList[pos] = logList
	time.Sleep(time.Second)
}

func addFile(file []string, fileLogs []string) []string {
	if len(file) <= 1 {
		fmt.Println("Write local url of the file")
	} else {
		var file = file[1]
		if fileutil.IsLocalFile(file) && fileutil.IsValidFileExtension(file) {

			if !fileLogsContainsFile(file, fileLogs) {
				filesIn := []string{file}
				for _, f := range filesIn {
					fileLogs = append(fileLogs, f)
				}
			} else {
				fmt.Println("This file was loaded before!")
			}
		}
	}
	return fileLogs
}

func fileLogsContainsFile(fileLog string, fileLogs []string) bool {
	var containsFile = false
	for _, v := range fileLogs {
		if v == fileLog {
			containsFile = true
			break
		}
	}
	return containsFile
}

func displayFileLogs(fileLog []string) {
	if len(fileLog) > 0 {
		fmt.Println("\nLoaded files:")
		for _, v := range fileLog {
			fmt.Println(v)
		}
		fmt.Println(" ")
	}
}

func title() {
	c := color.New(color.FgYellow)
	c.Println("_______________\n| Sherlog 1.0 |\n_______________")
	fmt.Println()
}
