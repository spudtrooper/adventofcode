package must

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/spudtrooper/adventofcode/common"
	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/must"
)

func chk(err error) {
	check.Err(err)
}

func Check(err error) {
	chk(err)
}

func ParseInt(s string, base, bits int) int64 {
	res, err := strconv.ParseInt(s, base, bits)
	chk(err)
	return res
}

func ReadLines(input string) []string {
	res, err := common.ReadLines(input)
	chk(err)
	return res
}

func ReadFile(input string, parse func(line string) (interface{}, error)) []interface{} {
	res, err := common.ReadFile(input, func(line string) (interface{}, error) {
		v, err := parse(line)
		chk(err)
		return v, nil
	})
	chk(err)
	return res
}

func Atoi(s string) int {
	res, err := strconv.Atoi(s)
	chk(err)
	return res
}

func ReadAllFile(input string) string {
	b, err := ioutil.ReadFile(input)
	chk(err)
	return string(b)
}

func SplitInts(s, sep string) []int {
	var res []int
	for _, s := range strings.Split(s, sep) {
		res = append(res, must.Atoi(s))
	}
	return res
}
