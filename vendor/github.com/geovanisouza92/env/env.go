package env

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var (
	ErrIgnoredLine      = errors.New("Ignored line")
	ErrCantSeparateLine = errors.New("")
)

func Load(filenames ...string) {
	filenames = filenamesOrDefault(filenames)

	for _, filename := range filenames {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			continue
		}
		loadFile(filename)
	}
}

func loadFile(filename string) error {
	envMap, err := readFile(filename)
	if err != nil {
		return err
	}

	for k, v := range envMap {
		if os.Getenv(k) == "" {
			os.Setenv(k, v)
		}
	}

	return nil
}

func readFile(filename string) (envMap map[string]string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	envMap = make(map[string]string)

	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	for _, l := range lines {
		if k, v, err := parseLine(l); err == nil {
			envMap[k] = v
		}
	}

	return
}

func parseLine(line string) (k, v string, err error) {
	if isIgnoredLine(line) {
		err = ErrIgnoredLine
		return
	}

	if strings.Contains(line, "#") {
		segments := strings.Split(line, "#")
		quotesAreOpen := false
		var segmentsToKeep []string
		for _, segment := range segments {
			if strings.Count(segment, "\"") == 1 || strings.Count(segment, "'") == 1 {
				if quotesAreOpen {
					quotesAreOpen = false
					segmentsToKeep = append(segmentsToKeep, segment)
				} else {
					quotesAreOpen = true
				}
			}

			if len(segmentsToKeep) == 0 || quotesAreOpen {
				segmentsToKeep = append(segmentsToKeep, segment)
			}
		}

		line = strings.Join(segmentsToKeep, "#")
	}

	s := strings.SplitN(line, "=", 2)

	if len(s) != 2 {
		s = strings.SplitN(line, ":", 2)
	}

	if len(s) != 2 {
		err = ErrCantSeparateLine
	}

	k = s[0]
	if strings.HasPrefix(k, "export") {
		k = strings.TrimPrefix(k, "export")
	}
	k = strings.Trim(k, " ")

	v = strings.Trim(s[1], " ")
	if strings.Count(v, "\"") == 2 || strings.Count(v, "'") == 2 {
		v = strings.Trim(v, "\"'")
		v = strings.Replace(v, "\\\"", "\"", -1)
		v = strings.Replace(v, "\\n", "\n", -1)
	}

	return
}

func isIgnoredLine(line string) bool {
	line = strings.Trim(line, " \n\t")
	return len(line) == 0 || strings.HasPrefix(line, "#")
}

func filenamesOrDefault(filenames []string) []string {
	if len(filenames) == 0 {
		return []string{".env"}
	}
	return filenames
}
