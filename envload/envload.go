package envload

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func init() {
	file, err := os.Open(".env")
	if err != nil {
		return
	}

	for key, value := range parseVars(file) {
		err := os.Setenv(key, value)
		if err != nil {
			return
		}
	}
}

func parseVars(r io.Reader) map[string]string {
	vars := make(map[string]string)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		pair := strings.SplitN(line, "=", 2)
		if len(pair) != 2 {
			continue
		}

		key := strings.TrimSpace(pair[0])
		value := strings.TrimSpace(pair[1])
		value = strings.Trim(value, "\"")

		vars[key] = value
	}

	return vars
}
