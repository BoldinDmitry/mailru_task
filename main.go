package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

// ReadInput splits input string by line
// It returns slice of splited strings
func ReadInput(in io.Reader) []string {
	scanner := bufio.NewScanner(in)
	var links []string
	for scanner.Scan() {
		links = append(links, scanner.Text())
	}
	return links
}

// IsValidURL checks is given string a valid URL
func IsValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// GetFile returns file data
func GetFile(path string) ([]byte, error) {
	document, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}
	return document, nil
}

// GetWebPage returns web page HTML-code
func GetWebPage(path string) ([]byte, error) {
	resp, err := http.Get(path)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	document, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return document, nil
}

// GetWordFrequency counts substring "word" in "document"
//
// Note: GetWordFrequency count substrings, not words
func GetWordFrequency(document, word []byte) int {
	return bytes.Count(document, word)
}

type WordsCountForPath struct {
	Path  string
	Count int
}

// CountWordFrequency returns struct WordsCountForPath
//
// Note: if there was any error, while getting web page
// or file "Path" field will be -1
func CountWordFrequency(path string, word []byte) WordsCountForPath {
	var document []byte
	var err error
	if IsValidURL(path) {
		document, err = GetWebPage(path)
	} else {
		document, err = GetFile(path)
	}
	if err != nil {
		return WordsCountForPath{
			Path:  path,
			Count: -1,
		}
	}

	return WordsCountForPath{
		Path:  path,
		Count: GetWordFrequency(document, word),
	}
}

// CountWordFrequencyGoroutine provides use of function
// CountWordFrequency using channels
func CountWordFrequencyGoroutine(tasksQueue chan string, word []byte,
	done chan WordsCountForPath) {
	for {
		path, isNewTask := <-tasksQueue
		if isNewTask {
			done <- CountWordFrequency(path, word)
		} else {
			return
		}
	}
}

// SolveTask reads all data from Stdin, creates goroutines
// then creates two channels(tasksQueue, countOfWords): for tasks and
// answers and fills first one, after that it reads all data from countOfWords
// channel and closes tasksQueue channel(for stopping all goroutines)
func SolveTask() {
	input := ReadInput(os.Stdin)

	wordToCount := []byte("Go")
	k := 5
	var numberOfWorkers int

	if len(input) > k {
		numberOfWorkers = k
	} else {
		numberOfWorkers = len(input)
	}

	tasksQueue := make(chan string)
	countOfWords := make(chan WordsCountForPath)

	for i := 0; i < numberOfWorkers; i++ {
		go CountWordFrequencyGoroutine(tasksQueue, wordToCount, countOfWords)
	}

	for _, path := range input {
		go func(path string) {
			tasksQueue <- path
		}(path)
	}

	var total int
	for i := 0; i < len(input); i++ {
		countForaPath := <-countOfWords

		if countForaPath.Count == -1 {
			fmt.Printf("Error while getting %s\n", countForaPath.Path)
		} else {
			fmt.Printf("Count for %s: %d\n", countForaPath.Path, countForaPath.Count)
			total += countForaPath.Count
		}
	}
	fmt.Printf("Total: %d\n", total)

	close(tasksQueue)
	time.Sleep(2 * time.Second)
}

func main() {
	SolveTask()
}
