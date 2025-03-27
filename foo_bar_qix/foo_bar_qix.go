package foobarqix

import "strconv"

const foo = "Foo"
const bar = "Bar"
const qix = "Qix"

func Compute(input string) string {
	v, err := strconv.Atoi(input)
	if err != nil {
		return input
	}

	var result string

	if v%3 == 0 {
		result += foo
	}
	if v%5 == 0 {
		result += bar
	}
	if v%7 == 0 {
		result += qix
	}

	for _, char := range input {
		if char == '3' {
			result += foo
		}
		if char == '5' {
			result += bar
		}
		if char == '7' {
			result += qix
		}
	}

	if result == "" {
		return input
	}

	return result
}
