package main

import "testing"

func TestAddFiles1(t *testing.T) {
	fileLogs := []string{}
	file := []string{"search", "C:\\server.log"}

	fileLogs = addFile(file, fileLogs)

	if len(fileLogs) != 1 {
		t.Error("Expected 1, got ", len(fileLogs))
	}
}

func TestAddFiles2(t *testing.T) {
	fileLogs := []string{}
	file := []string{"C:\\server.log"}

	fileLogs = addFile(file, fileLogs)

	if len(fileLogs) != 0 {
		t.Error("Expected 0, got ", len(fileLogs))
	}
}

func TestAddFiles3(t *testing.T) {
	fileLogs := []string{"C:\\server.log"}
	file := []string{"search", "C:\\server.log"}

	fileLogs = addFile(file, fileLogs)

	if len(fileLogs) != 1 {
		t.Error("Expected 1, got ", len(fileLogs))
	}
}

func TestFileLogsContainsFile1(t *testing.T) {
	fileLogs := []string{"C:\\server.log"}

	ok := fileLogsContainsFile("C:\\server.log", fileLogs)
	if !ok {
		t.Error("Expected file 'C:\\server.log'")
	}

	ok = fileLogsContainsFile("C:\\server2.log", fileLogs)
	if ok {
		t.Error("Non Expected file 'C:\\server2.log'")
	}
}

func TestSearchFiles1(t *testing.T) {
	fileLogs := []string{"C:\\server.log"}
	words := []string{"search", "command"}
	searchFiles(words, fileLogs, false)
	searchFiles(words, fileLogs, true)

	words = []string{"command"}
	searchFiles(words, fileLogs, false)

	fileLogs = []string{}
	words = []string{}
	searchFiles(words, fileLogs, false)
}

func TestDisplayFileLogs(t *testing.T) {
	fileLogs := []string{"C:\\server.log"}
	displayFileLogs(fileLogs)
}
