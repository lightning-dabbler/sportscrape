package util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// re is the regex compilation of one or more spaces
var re = regexp.MustCompile(`\s+`)

// StrFormat runs a string replace on a formatted string
// For example StrFormat("{foo} bar {baz}","foo","test","baz","result") --> "test bar result"
func StrFormat(format string, args ...string) (string, error) {
	argsLen := len(args)
	if argsLen%2 != 0 {
		return "", fmt.Errorf("Invalid number of arguments to replace: %d. Number of arguments must be even", argsLen)
	}
	for i, v := range args {
		if i%2 == 0 {
			args[i] = "{" + v + "}"
		}
	}
	r := strings.NewReplacer(args...)
	return r.Replace(format), nil
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
	val, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("Could not convert %s to int: %w", str, err)
	}
	return val, nil
}

// TextToInt32
func TextToInt32(str string) (int32, error) {
	val, err := TextToInt(str)
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

// TextToInt64 attempts to convert a string to an int64
// Returns an 64-bit integer and error (nil if the transformation was successful)
func TextToInt64(str string) (int64, error) {
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Could not convert %s to int64: %w", str, err)
	}
	return val, nil
}

// TextToFloat64 attempts to convert a string to a float64
// Returns an 64-bit float and error (nil if the transformation was successful)
func TextToFloat64(str string) (float64, error) {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, fmt.Errorf("Could not convert %s to float64: %w", str, err)
	}
	return val, nil
}

// TextToFloat32 attempts to convert a string to a float32
// Returns an 32-bit float and error (nil if the transformation was successful)
func TextToFloat32(str string) (float32, error) {
	val, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, fmt.Errorf("Could not convert %s to float32: %w", str, err)
	}
	return float32(val), nil
}
