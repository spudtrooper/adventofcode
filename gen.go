package main

import (
	"flag"

	"github.com/go-errors/errors"
	"github.com/spudtrooper/adventofcode/common/must"
	"github.com/spudtrooper/adventofcode/gen"
)

var (
	year = flag.Int("year", 0, "year to generate")
	day  = flag.Int("day", 0, "day to generate")
)

func realMain() error {
	if *year == 0 {
		return errors.Errorf("--year required")
	}
	if *day == 0 {
		return errors.Errorf("--day required")
	}
	if err := gen.Main(*year, *day); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	must.Check(realMain())
}
