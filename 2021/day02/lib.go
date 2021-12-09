package day02

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/spudtrooper/adventofcode/common"
)

type cmd struct {
	move string
	val  int
}

func parseCmd(line string) (interface{}, error) {
	parts := strings.Split(line, " ")
	move := parts[0]
	val, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	return cmd{move: move, val: val}, nil
}

type sub struct {
	hPos, depth int
}

func (s *sub) Move(input string) (int, error) {
	cmds, err := common.ReadFile(input, parseCmd)
	if err != nil {
		return 0, err
	}

	for _, c := range cmds {
		switch c := c.(cmd); c.move {
		case "forward":
			s.hPos += c.val
		case "up":
			s.depth -= c.val
		case "down":
			s.depth += c.val
		default:
			return 0, errors.Errorf("Move: invalid move: %+v", c.move)
		}
	}

	return s.hPos * s.depth, nil
}

type sub2 struct {
	hPos, depth, aim int
}

func (s *sub2) Move(input string) (int, error) {
	cmds, err := common.ReadFile(input, parseCmd)
	if err != nil {
		return 0, err
	}

	for _, c := range cmds {
		switch c := c.(cmd); c.move {
		case "forward":
			s.hPos += c.val
			s.depth += s.aim * c.val
		case "up":
			s.aim -= c.val
		case "down":
			s.aim += c.val
		default:
			return 0, errors.Errorf("Move: invalid move: %s ", c.move)
		}
	}

	return s.hPos * s.depth, nil
}
