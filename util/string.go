package util

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// re is the regex compilation of one or more spaces
var re = regexp.MustCompile(`\s+`)

// StrFormat runs a string replace on a formatted string
// For example StrFormat("{foo} bar {baz}","foo","test","baz","result") --> "test bar result"
func StrFormat(format string, args ...string) string {
	argsLen := len(args)
	if argsLen%2 != 0 {
		log.Fatalf("Invalid number of arguments to replace: %d\n", argsLen)
	}
	for i, v := range args {
		if i%2 == 0 {
			args[i] = "{" + v + "}"
		}
	}
	r := strings.NewReplacer(args...)
	return r.Replace(format)
}

// CleanTextDatum removes leading and trailing spaces and replaces multiple succeeding spaces with one space
// Returns a copy of the string with its replacements
func CleanTextDatum(str string) string {
	trimmedStr := strings.TrimSpace(str)
	return re.ReplaceAllString(trimmedStr, " ")
}

// TextToInt attempts to convert a string to an int
// Returns an integer and error (nil if the transformation was successful)
func TextToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// TextToInt64 attempts to convert a string to an int64
// Returns an 64-bit integer and error (nil if the transformation was successful)
func TextToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// TextToFloat64 attempts to convert a string to a float64
// Returns an 64-bit float and error (nil if the transformation was successful)
func TextToFloat64(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}
