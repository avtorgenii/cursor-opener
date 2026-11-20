package projects

import (
	"os"
	"strings"
)

func List(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	var names []string
	for _, e := range entries {
		if e.IsDir() && !strings.HasPrefix(e.Name(), ".") {
			names = append(names, e.Name())
		}
		
	}

	return names, nil
}
