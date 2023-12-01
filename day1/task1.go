package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type posVal struct {
	position int
	value    int
}

func NewPosVal(position int, value int) *posVal {
	p := posVal{
		position: position,
		value:    value,
	}
	return &p
}

func GetLastNumericString(line string) posVal {
	digitWords := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var lastValue posVal = *NewPosVal(-1, -1)

	for val := 0; val < len(digitWords); val++ {
		var word = digitWords[val]
		var index = strings.LastIndex(line, word)
		if index >= 0 && lastValue.position < 0 {
			lastValue.position = index
			lastValue.value = val
		} else if index >= 0 && lastValue.position >= 0 && index > lastValue.position {
			lastValue.position = index
			lastValue.value = val
		}
	}
	return lastValue
}

func GetFirstNumericString(line string) posVal {
	digitWords := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var firstValue posVal = *NewPosVal(-1, -1)

	for val := 0; val < len(digitWords); val++ {
		var word = digitWords[val]
		var index = strings.Index(line, word)

		if index >= 0 && firstValue.position < 0 {
			firstValue.position = index
			firstValue.value = val
		} else if index >= 0 && firstValue.position >= 0 && index < firstValue.position {
			firstValue.position = index
			firstValue.value = val
		}
	}
	return firstValue
}

func GetFirstdigit(line string) posVal {
	var firstValue posVal = *NewPosVal(-1, -1)

	for pos, rune := range line {
		if unicode.IsDigit(rune) {
			firstValue.value = int(rune - '0')
			firstValue.position = pos
			break
		}
	}
	if firstValue.value < 0 {
		errors.New("unable to parse line: " + line)
	}
	return firstValue
}

func GetLastdigit(line string) posVal {
	var lastValue posVal = *NewPosVal(-1, -1)

	for pos, rune := range line {
		if unicode.IsDigit(rune) {
			lastValue.value = int(rune - '0')
			lastValue.position = pos
		}
	}
	if lastValue.value < 0 {
		errors.New("unable to parse line: " + line)
	}
	return lastValue
}

func ExtractValueFromLine(line string) int {
	var firstDigitValue = GetFirstdigit(line)
	var lastDigitValue = GetLastdigit(line)

	var firstTextValue = GetFirstNumericString(line)
	var lastTextValue = GetLastNumericString(line)

	var firstValue int = 0
	var lastValue int = 0

	// first value
	if firstDigitValue.position >= 0 && firstTextValue.position >= 0 {
		if firstDigitValue.position < firstTextValue.position {
			firstValue = firstDigitValue.value
		} else {
			firstValue = firstTextValue.value
		}
	}

	if firstDigitValue.position < 0 {
		firstValue = firstTextValue.value
	}
	if firstTextValue.position < 0 {
		firstValue = firstDigitValue.value
	}

	//Last value
	if lastDigitValue.position >= 0 && lastTextValue.position >= 0 {
		if lastDigitValue.position > lastTextValue.position {
			lastValue = lastDigitValue.value
		} else {
			lastValue = lastTextValue.value
		}
	}

	if lastDigitValue.position < 0 {
		lastValue = lastTextValue.value
	}
	if lastTextValue.position < 0 {
		lastValue = lastDigitValue.value
	}

	var lineValue int = firstValue*10 + lastValue
	return lineValue
}

func main() {

	file, err := os.Open("task1_input.txt")
	if err != nil {
		fmt.Println("Error opening file.")
	}

	defer file.Close()

	var line string
	scanner := bufio.NewScanner(file)
	var sumOfAllValues int = 0

	for scanner.Scan() {
		line = scanner.Text()
		var lineValue = ExtractValueFromLine(line)
		sumOfAllValues += lineValue

	}

	fmt.Println("Sum of it all is ", sumOfAllValues)

	file.Close()
}
