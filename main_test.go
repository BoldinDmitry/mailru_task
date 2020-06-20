package main

import (
	"fmt"
	"strings"
	"testing"
)

// This function was written to prevent of using assert (non-standard library)
func EqualStringsSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestReadInput(t *testing.T) {
	type readInputTestCase struct {
		inputString string
		outputArray []string
	}

	testCases := []readInputTestCase{
		{
			inputString: "https://golang.org\n/etc/passwd\nhttps://golang.org\nhttps://golang.org",
			outputArray: []string{
				"https://golang.org",
				"/etc/passwd",
				"https://golang.org",
				"https://golang.org",
			},
		},
		{
			inputString: "https://golang.org",
			outputArray: []string{
				"https://golang.org",
			},
		},
		{
			inputString: "http://localhost:8080\n",
			outputArray: []string{
				"http://localhost:8080",
			},
		},
	}

	for i, testCase := range testCases {
		inputReader := strings.NewReader(testCase.inputString)
		outputArray := ReadInput(inputReader)

		if !EqualStringsSlices(testCase.outputArray, outputArray) {
			message := fmt.Sprintf("error in test case number: %d", i)
			t.Error(message)
		}
	}
}

func TestIsValidURL(t *testing.T) {
	type readInputTestCase struct {
		inputURL string
		answer   bool
	}

	testCases := []readInputTestCase{
		{
			inputURL: "https://golang.org",
			answer:   true,
		},
		{
			inputURL: "http://localhost:8080",
			answer:   true,
		},
		{
			inputURL: "/etc/passwd",
			answer:   false,
		},
		{
			inputURL: "NOT URL",
			answer:   false,
		},
	}

	for i, testCase := range testCases {
		funcAnswer := IsValidURL(testCase.inputURL)

		if testCase.answer != funcAnswer {
			message := fmt.Sprintf("error in test case number: %d", i)
			t.Error(message)
		}
	}
}

func TestGetWordFrequency(t *testing.T) {
	type readInputTestCase struct {
		inputDocument []byte
		word          []byte
		answer        int
	}

	testCases := []readInputTestCase{
		{
			inputDocument: []byte("Let's code Go!"),
			word:          []byte("Go"),
			answer:        1,
		},
		{
			inputDocument: []byte("There is no g-word"),
			word:          []byte("Go"),
			answer:        0,
		},
		// Answer is 2, because "Google" contains substring "Go"
		{
			inputDocument: []byte("Go is language developed by Google"),
			word:          []byte("Go"),
			answer:        2,
		},
	}

	for i, testCase := range testCases {
		funcAnswer := GetWordFrequency(testCase.inputDocument, testCase.word)

		if testCase.answer != funcAnswer {
			message := fmt.Sprintf("error in test case number: %d", i)
			t.Error(message)
		}
	}
}
