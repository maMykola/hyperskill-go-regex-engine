package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := scanner.Text()

	inputs := strings.Split(data, "|")

	fmt.Println(match(inputs[0], inputs[1]))
}

func match(regex string, input string) bool {
	var matchEnd = false

	if len(regex) == 0 {
		return true
	}

	if len(regex) > 0 && regex[len(regex)-1] == '$' {
		matchEnd = true
		regex = regex[:len(regex)-1]
	}

	if len(regex) > 0 && regex[0] == '^' {
		return regexMatch(regex[1:], input, matchEnd)
	}

	for {
		if regexMatch(regex, input, matchEnd) {
			return true
		}

		if len(input) == 0 {
			return false
		}

		input = input[1:]
	}
}

func regexMatch(regex string, input string, matchEnd bool) bool {
	if len(regex) == 0 {
		return !matchEnd || len(input) == 0
	}

	if len(input) == 0 {
		return false
	}

	if regex[0] == '\\' {
		regex = regex[1:]
	}

	repeater := getRepeater(regex)

	if regex[0] != '.' && regex[0] != input[0] {
		if repeater == '?' || repeater == '*' {
			return regexMatch(regex[2:], input, matchEnd)
		} else {
			return false
		}
	}

	switch repeater {
	case '?':
		return regexMatch(regex[2:], input[1:], matchEnd)
	case '+':
		return regexMatch(replacePlusRepeater(regex), input[1:], matchEnd)
	case '*':
		return regexMatch(regex, input[1:], matchEnd) || regexMatch(regex[2:], input, matchEnd)
	default:
		return regexMatch(regex[1:], input[1:], matchEnd)
	}
}

func getRepeater(regex string) byte {
	if len(regex) < 2 {
		return 0
	}

	return regex[1]
}

func replacePlusRepeater(regex string) string {
	newRegex := make([]byte, 0, len(regex))
	newRegex = append(newRegex, regex[0])
	newRegex = append(newRegex, '*')
	newRegex = append(newRegex, regex[2:]...)
	return string(newRegex)
}
