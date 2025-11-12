package libvalidator

import (
	"os"
	"path/filepath"
	"strings"
)

func FindExec(name string) (string, bool) {
	if filepath.IsAbs(name) {
		_, err := os.Stat(name)
		return name, err == nil
	}

	for dir := range strings.SplitSeq(os.Getenv("PATH"), ":") {
		file_path := filepath.Join(dir, name)

		_, err := os.Stat(file_path)
		if err == nil {
			return file_path, true
		}
	}

	return name, false
}
