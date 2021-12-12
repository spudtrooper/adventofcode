package day03

import (
	"fmt"
	"strconv"

	"github.com/spudtrooper/adventofcode/common"
)

type result struct {
	gamma, epsilon int
}

func (r *result) Multiply() int {
	return r.gamma * r.epsilon
}

func (r *result) String() string {
	return fmt.Sprintf("gamma=%b (%d), epsilon=%b (%d), mul=%d",
		r.gamma, r.gamma, r.epsilon, r.epsilon, r.Multiply())
}

func Part1(input string) (*result, error) {
	codes, err := common.ReadLines(input)
	if err != nil {
		return nil, err
	}

	var gamma, epsilon int
	for i := range codes[0] {
		var diff int
		for _, c := range codes {
			if string(c[i]) == "1" {
				diff++
			} else {
				diff--
			}
		}
		gamma <<= 1
		epsilon <<= 1
		if diff > 0 {
			gamma++
		} else {
			epsilon++
		}
	}

	res := &result{
		gamma:   gamma,
		epsilon: epsilon,
	}
	return res, nil
}

type part2result struct {
	ogr, csr int
}

func (r *part2result) Multiply() int {
	return r.ogr * r.csr
}

func (r *part2result) String() string {
	return fmt.Sprintf("ogr=%b (%d), csr=%b (%d), mul=%d",
		r.ogr, r.ogr, r.csr, r.csr, r.Multiply())
}

func Part2(input string) (*part2result, error) {
	codes, err := common.ReadLines(input)
	if err != nil {
		return nil, err
	}

	findDiffAndSeparate := func(i int, arr []string) (diff int, ones, zeros []string) {
		for _, c := range arr {
			if string(c[i]) == "1" {
				diff++
				ones = append(ones, c)
			} else {
				diff--
				zeros = append(zeros, c)
			}
		}
		if diff == 0 {
			diff = 1
		}
		return
	}

	ogrs, csrs := codes, codes
	var ogr, csr int
	for i := range codes[0] {
		if diff, ones, zeros := findDiffAndSeparate(i, ogrs); diff > 0 {
			ogrs = ones
		} else {
			ogrs = zeros
		}
		if diff, ones, zeros := findDiffAndSeparate(i, csrs); diff > 0 {
			csrs = zeros
		} else {
			csrs = ones
		}

		if len(ogrs) == 1 {
			b, err := strconv.ParseInt(ogrs[0], 2, 32)
			if err != nil {
				return nil, err
			}
			ogr = int(b)
		}
		if len(csrs) == 1 {
			b, err := strconv.ParseInt(csrs[0], 2, 32)
			if err != nil {
				return nil, err
			}
			csr = int(b)
		}
		if ogr != 0 && csr != 0 {
			break
		}
	}

	res := &part2result{
		ogr: ogr,
		csr: csr,
	}
	return res, nil
}
