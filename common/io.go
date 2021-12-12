package common

import (
	"io/ioutil"
	"strings"
)

// ReadFile reads a generic array of items. This is slow but easier than using channels.
// parse is called on each line of the input
func ReadFile(input string, parse func(line string) (interface{}, error)) ([]interface{}, error) {
	b, err := ioutil.ReadFile(input)
	if err != nil {
		return nil, err
	}
	res := []interface{}{}
	if len(b) > 0 {
		for _, line := range strings.Split(string(b), "\n") {
			t, err := parse(line)
			if err != nil {
				return nil, err
			}
			res = append(res, t)
		}
	}
	return res, nil
}

// ReadLines reads a list of strings
func ReadLines(input string) ([]string, error) {
	b, err := ioutil.ReadFile(input)
	if err != nil {
		return nil, err
	}
	if len(b) == 0 {
		return []string{}, nil
	}
	return strings.Split(string(b), "\n"), nil
}

func Split(input, sep string, parse func(line string) (interface{}, error)) ([]interface{}, error) {
	res := []interface{}{}
	for _, s := range strings.Split(input, sep) {
		t, err := parse(s)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}
