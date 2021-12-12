package day08

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/spudtrooper/adventofcode/common/must"
)

func Part1(input string) int {
	var uniqueValues int
	for _, line := range must.ReadLines(input) {
		parts := strings.Split(line, "|")
		outputValues := strings.Split(strings.TrimSpace(parts[1]), " ")
		for _, s := range outputValues {
			switch len(s) {
			case 2 /* '1' */, 4 /* '4' */, 3 /* '7' */, 7 /* '8' */ :
				uniqueValues++

			}
		}
	}
	return uniqueValues
}

// https://stackoverflow.com/questions/22688651/golang-how-to-sort-string-or-byte
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// https://yourbasic.org/golang/generate-permutation-slice-string/
// permString calls f with each permutation of a.
func permString(a string, f func(string)) {
	permStringRec([]rune(a), f, 0)
}

// Permute the values at index i to len(a)-1.
func permStringRec(a []rune, f func(string), i int) {
	if i > len(a) {
		s := string(a)
		f(s)
		return
	}
	permStringRec(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permStringRec(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func permStrings(a []string, f func([]string)) {
	permStringsRec(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func permStringsRec(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}
	permStringsRec(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permStringsRec(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func Part2(input string) int {
	/*
		  0:      1:      2:      3:      4:
		 aaaa    ....    aaaa    aaaa    ....
		b    c  .    c  .    c  .    c  b    c
		b    c  .    c  .    c  .    c  b    c
		 ....    ....    dddd    dddd    dddd
		e    f  .    f  e    .  .    f  .    f
		e    f  .    f  e    .  .    f  .    f
		 gggg    ....    gggg    gggg    ....

		  5:      6:      7:      8:      9:
		 aaaa    aaaa    aaaa    aaaa    aaaa
		b    .  b    .  .    c  b    c  b    c
		b    .  b    .  .    c  b    c  b    c
		 dddd    dddd    ....    dddd    dddd
		.    f  e    f  .    f  e    f  .    f
		.    f  e    f  .    f  e    f  .    f
		 gggg    gggg    ....    gggg    gggg
	*/

	/*
			Assignments:

			  0000
			 1    2
			 1    2
			  3333
		     4    5
			 4    5
			  6666
	*/

	createPatFromOutput := func(perm string, indices ...int) string {
		var buf bytes.Buffer
		for _, i := range indices {
			buf.WriteByte(perm[i])
		}
		return sortString(buf.String())
	}

	var outputValueSum int
	for _, line := range must.ReadLines(input) {
		parts := strings.Split(line, "|")
		outputs := strings.Split(strings.TrimSpace(parts[0]), " ")
		signals := strings.Split(strings.TrimSpace(parts[1]), " ")

		lengthToPats := map[int][]string{}
		for _, p := range outputs {
			lengthToPats[len(p)] = append(lengthToPats[len(p)], sortString(p))
		}

		var outputPatterns map[string]int

		// 8
		permString("abcdefg", func(perm string) {
			outputPats := map[int]string{8: perm}

			// 4
			{
				patFromOutput := createPatFromOutput(perm, 1, 2, 3, 5)
				if pat := lengthToPats[4][0]; pat != patFromOutput {
					return
				}
				outputPats[4] = patFromOutput
			}
			// 7
			{
				patFromOutput := createPatFromOutput(perm, 0, 2, 5)
				if pat := lengthToPats[3][0]; pat != patFromOutput {
					return
				}
				outputPats[7] = patFromOutput
			}
			// 1
			{
				patFromOutput := createPatFromOutput(perm, 2, 5)
				if pat := lengthToPats[2][0]; pat != patFromOutput {
					return
				}
				outputPats[1] = patFromOutput
			}

			// 0, 6, 9
			permStrings(lengthToPats[6], func(pats []string) {
				// 0
				{
					patFromOutput := createPatFromOutput(perm, 0, 1, 2, 4, 5, 6)
					if pat := pats[0]; pat != patFromOutput {
						return
					}
					outputPats[0] = patFromOutput
				}
				// 6
				{
					patFromOutput := createPatFromOutput(perm, 0, 1, 3, 4, 5, 6)
					if pat := pats[1]; pat != patFromOutput {
						return
					}
					outputPats[6] = patFromOutput
				}
				// 9
				{
					patFromOutput := createPatFromOutput(perm, 0, 1, 2, 3, 5, 6)
					if pat := pats[2]; pat != patFromOutput {
						return
					}
					outputPats[9] = patFromOutput
				}

				// 2, 3, 5
				permStrings(lengthToPats[5], func(pats []string) {
					// 2
					{
						patFromOutput := createPatFromOutput(perm, 0, 2, 3, 4, 6)
						if pat := pats[0]; pat != patFromOutput {
							return
						}
						outputPats[2] = patFromOutput
					}
					// 3
					{
						patFromOutput := createPatFromOutput(perm, 0, 2, 3, 5, 6)
						if pat := pats[1]; pat != patFromOutput {
							return
						}
						outputPats[3] = patFromOutput
					}
					// 5
					{
						patFromOutput := createPatFromOutput(perm, 0, 1, 3, 5, 6)
						if pat := pats[2]; pat != patFromOutput {
							return
						}
						outputPats[5] = patFromOutput
					}
				})
			})

			if len(outputPats) < 9 {
				return
			}

			if outputPatterns != nil {
				panic("outputPatterns should be nil")
			}

			outputPatterns = map[string]int{}
			for x, o := range outputPats {
				outputPatterns[sortString(o)] = x
			}
		})

		if outputPatterns == nil {
			panic("outputPatterns is nil")
		}

		var outputValue int
		for _, v := range signals {
			n, ok := outputPatterns[sortString(v)]
			if !ok {
				panic(fmt.Sprintf("line=%q cannot find entry for %s outputPatterns=%v", line, v, outputPatterns))
			}
			outputValue = outputValue*10 + n
		}
		outputValueSum += outputValue
	}

	return outputValueSum
}
