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

	"github.com/spudtrooper/goutil/io"
)

func Main(year, day int) error {
	pkg := fmt.Sprintf("day%02d", day)
	dir, err := io.MkdirAll(fmt.Sprintf("%d", year), pkg)
	if err != nil {
		return err
	}

	lib, err := writeFile(`	
package {{.Pkg}}

func Part1(input string) int {
	// TODO
	return -1
}

func Part2(input string) int {
	// TODO
	return -1
}
`, struct {
		Pkg string
	}{
		Pkg: pkg,
	}, dir, "lib.go")
	if err != nil {
		return err
	}

	libTest, err := writeFile(`
package {{.Pkg}}

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
		{
			name:  "part1",
			input: "testdata/input.txt",
			want:  -1, // TODO
		},
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
	}, dir, "lib_test.go")
	if err != nil {
		return err
	}

	testdataDir, err := io.MkdirAll(dir, "testdata")
	if err != nil {
		return err
	}

	if err := touch(testdataDir, "input.txt"); err != nil {
		return err
	}
	if err := touch(testdataDir, "testinput.txt"); err != nil {
		return err
	}

	if err := run("go", "fmt", lib, libTest); err != nil {
		return err
	}

	if err := run("go", "test", lib, libTest); err != nil {
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
func run(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
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
