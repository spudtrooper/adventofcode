package gen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/io"
	"github.com/spudtrooper/goutil/task"
)

func Main(year, day int, mainOpts ...Option) error {
	opts := makeOptionImpl(mainOpts...)

	pkg := fmt.Sprintf("day%02d", day)

	if !opts.force {
		dayDir := path.Join(fmt.Sprintf("%d", year), pkg)
		if io.FileExists(dayDir) {
			return errors.Errorf("%s exists, pass --force to overwrite it", dayDir)
		}
	}

	dayDir, err := io.MkdirAll(fmt.Sprintf("%d", year), pkg)
	if err != nil {
		return err
	}
	libDir, err := io.MkdirAll(dayDir, "lib")
	if err != nil {
		return err
	}

	lib, err := writeFile(`	
package lib

import (
	"log"

	"github.com/spudtrooper/adventofcode/common/must"
)


func Part1(input string) int {
	for _, line := range must.ReadLines(input) {
		// TODO
		if false {
			log.Println(line)
		}
	}
	return -1
}

func Part2(input string) int {
	for _, line := range must.ReadLines(input) {
		// TODO
		if false {
			log.Println(line)
		}
	}
	return -1
}
`, struct {
		Pkg string
	}{
		Pkg: pkg,
	}, libDir, "lib.go")
	if err != nil {
		return err
	}

	libTest, err := writeFile(`
package lib

import "testing"

func TestPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "testinput",
			input: "testdata/testinput.txt",
			want:  -1, // TODO
		},
		// {
		// 	name:  "part1",
		//	input: "testdata/input.txt",
		//	want:  -1, // TODO
		// },
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Part1(tc.input); want != got {
				t.Errorf("Part1: want(%d) != got(%d)", want, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "testinput",
			input: "testdata/testinput.txt",
			want:  -1, // TODO
		},
		// {
		// 	name:  "part2",
		//	input: "testdata/input.txt",
		//	want:  -1, // TODO
		// },
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Part2(tc.input); want != got {
				t.Errorf("Part2: want(%d) != got(%d)", want, got)
			}
		})
	}
}	
`, struct {
		Pkg string
	}{
		Pkg: pkg,
	}, libDir, "lib_test.go")
	if err != nil {
		return err
	}

	writeMain := func(n int) (string, error) {
		return writeFile(`	
package main

import (
	"flag"
	"fmt"

	{{.Pkg}} "github.com/spudtrooper/adventofcode/{{.Year}}/{{.Pkg}}/lib"
)

var (
	input = flag.String("input", "{{.Year}}/{{.Pkg}}/lib/testdata/testinput.txt", "test input")
)

func main() {
	flag.Parse()
	fmt.Printf("{{.Pkg}} Part{{.N}}: %d\n", {{.Pkg}}.Part{{.N}}(*input))
}`, struct {
			Pkg  string
			Year int
			N    int
		}{
			Pkg:  pkg,
			Year: year,
			N:    n,
		}, dayDir, fmt.Sprintf("part%d.go", n))
	}

	testdataDir, err := io.MkdirAll(libDir, "testdata")
	if err != nil {
		return err
	}

	mainPart1, err := writeMain(1)
	if err != nil {
		return err
	}
	mainPart2, err := writeMain(2)
	if err != nil {
		return err
	}

	tb := task.MakeBuilder(task.Color(color.New(color.FgYellow)))

	tb.Add("creating input.txt", touchFn(testdataDir, "input.txt"))
	tb.Add("creating testinput.txt", touchFn(testdataDir, "testinput.txt"))

	tb.Add("formatting test files", runFn("go", "fmt", lib, libTest))
	tb.Add("formatting main files", runFn("go", "fmt", mainPart1, mainPart2))

	tb.Add("testing test files", runFn("go", "test", lib, libTest))

	tb.Add("testing main part1", runFn("go", "run", mainPart1))
	tb.Add("testing main part2", runFn("go", "run", mainPart2))
	tb.Add("done", func() error { return nil })

	if err := tb.Build().Go(); err != nil {
		return err
	}

	return nil
}

func writeFile(t string, data interface{}, dir string, outFileName string) (string, error) {
	b, err := renderTemplate(t, outFileName, data)
	if err != nil {
		return "", err
	}
	outFile := path.Join(dir, outFileName)
	if err := ioutil.WriteFile(outFile, b, 7055); err != nil {
		return "", err
	}

	log.Printf("wrote to %s", outFile)
	return outFile, nil
}

func touch(dir, outFileName string) error {
	outFile := path.Join(dir, outFileName)
	if err := ioutil.WriteFile(outFile, []byte{}, 7055); err != nil {
		return err
	}
	log.Printf("touched %s", outFile)
	return nil
}

func touchFn(dir, outFileName string) func() error {
	return func() error {
		return touch(dir, outFileName)
	}
}

func run(command string, args ...string) error {
	log.Printf("running %s %s", command, strings.Join(args, " "))
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func runFn(command string, args ...string) func() error {
	return func() error {
		return run(command, args...)
	}
}

func renderTemplate(t string, name string, data interface{}) ([]byte, error) {
	tmpl, err := template.New(name).Parse(strings.TrimSpace(t))
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
