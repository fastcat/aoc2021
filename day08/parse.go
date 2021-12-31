package day08

import "fmt"

func ParseEntry(line string) (Entry, error) {
	var patternStrings [10]string
	var outputStrings [4]string
	n, err := fmt.Sscanf(
		line,
		"%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
		&patternStrings[0], &patternStrings[1], &patternStrings[2], &patternStrings[3], &patternStrings[4],
		&patternStrings[5], &patternStrings[6], &patternStrings[7], &patternStrings[8], &patternStrings[9],
		&outputStrings[0], &outputStrings[1], &outputStrings[2], &outputStrings[3],
	)
	if err != nil {
		return Entry{}, err
	}
	if n != 14 {
		return Entry{}, fmt.Errorf("parse error, only got %d", n)
	}
	var ret Entry
	for i, s := range patternStrings {
		ret.Patterns[i], err = ParseValue(s)
		if err != nil {
			return ret, fmt.Errorf("invalid pattern '%s': %w", s, err)
		}
	}
	for i, s := range outputStrings {
		ret.Outputs[i], err = ParseValue(s)
		if err != nil {
			return ret, fmt.Errorf("invalud output '%s': %w", s, err)
		}
	}
	return ret, nil
}

func ParseEntries(input []string) ([]Entry, error) {
	ret := make([]Entry, len(input))
	var err error
	for i, s := range input {
		ret[i], err = ParseEntry(s)
		if err != nil {
			return ret, err
		}
	}
	return ret, nil
}

func ParseValue(pattern string) (Value, error) {
	var ret Value
	for _, r := range pattern {
		if r < 'a' || r > 'g' {
			return ret, fmt.Errorf("invalid item '%c'", r)
		}
		ret |= 1 << Value(r-'a')
	}
	return ret, nil
}
