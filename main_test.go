package main

import (
	"gopkg.in/go-playground/assert.v1"
	"strings"
	"testing"
)

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

	for _, testCase := range testCases {
		inputReader := strings.NewReader(testCase.inputString)
		outputArray := ReadInput(inputReader)

		assert.Equal(t, testCase.outputArray, outputArray)
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

	for _, testCase := range testCases {
		funcAnswer := IsValidURL(testCase.inputURL)

		assert.Equal(t, testCase.answer, funcAnswer)
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

	for _, testCase := range testCases {
		funcAnswer := GetWordFrequency(testCase.inputDocument, testCase.word)

		assert.Equal(t, testCase.answer, funcAnswer)
	}
}