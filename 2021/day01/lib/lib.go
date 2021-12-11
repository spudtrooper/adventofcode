package day01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spudtrooper/adventofcode/common"
)

type incType string

const (
	increased incType = "increased"
	decreased incType = "decreased"
	noChange  incType = "no change"
	na        incType = "(N/A - no previous measurement)"
)

type outputLine struct {
	level int
	inc   incType
}

type output []outputLine

func (c output) String() string {
	var lines []string
	for _, o := range c {
		lines = append(lines, fmt.Sprintf("%d %s", o.level, o.inc))
	}
	return strings.Join(lines, "\n")
}

func (c output) Increased() int {
	var res int
	for _, o := range c {
		if o.inc == increased {
			res++
		}
	}
	return res
}

func CountLevels(input string) (*output, error) {
	levels, err := common.ReadFile(input, func(s string) (interface{}, error) {
		return strconv.Atoi(s)
	})
	if err != nil {
		return nil, err
	}

	var out output
	for i, level := range levels {
		level := level.(int)
		inc := na
		if i > 0 {
			if last := levels[i-1].(int); level > last {
				inc = increased
			} else if level == last {
				inc = noChange
			} else {
				inc = decreased
			}
		}
		out = append(out, outputLine{
			level: level,
			inc:   inc,
		})
	}

	return &out, nil
}

func CountLevelsGrouped(input string, clusterSize int) (*output, error) {
	levels, err := common.ReadFile(input, func(s string) (interface{}, error) {
		return strconv.Atoi(s)
	})
	if err != nil {
		return nil, err
	}

	var groupedLevels []int
	for i := range levels {
		level := 0
		for j := 0; j < clusterSize && j+i < len(levels); j++ {
			level += levels[i+j].(int)
		}
		groupedLevels = append(groupedLevels, level)
	}

	var out output
	for i, level := range groupedLevels {
		inc := na
		if i > 0 {
			if last := groupedLevels[i-1]; level > last {
				inc = increased
			} else if level == last {
				inc = noChange
			} else {
				inc = decreased
			}
		}
		out = append(out, outputLine{
			level: level,
			inc:   inc,
		})
	}

	return &out, nil
}
