package internal

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Search(keyword, path string) {
	search(keyword, path)
}

func search(keyword, path string) {
	if !isValidPath(path) {
		return
	}
	// check if path is file or directory
	if isFile(path) {
		// if path is file, then search keyword in file with goroutine
		go printKeywordInFile(keyword, path)
	} else {
		// if path is directory, perform
		// for each elements in directory, search keyword in file or directory
		paths := getPathsInDirectory(path)
		for _, path := range paths {
			search(keyword, path)
		}
	}
}

func isFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	if fileInfo.IsDir() {
		return false
	}

	return true
}

func printKeywordInFile(keyword, path string) {
	readFile, err := os.Open(path)
	defer readFile.Close()
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	line := 1
	for fileScanner.Scan() {
		if strings.Contains(fileScanner.Text(), keyword) {
			println("File: " + path + " contains keyword: " + keyword + " in line: " + strconv.Itoa(line))
		}
		line++
	}
}

func getPathsInDirectory(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	paths := make([]string, len(files))
	for _, file := range files {
		path := path + "/" + file.Name()
		paths = append(paths, path)
	}

	return paths
}

func isValidPath(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// println("Path: " + path + " does not exist")
		return false
	}

	return true
}
