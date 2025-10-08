package utils

import "os"

func WriteFile(file, content string) error {
	return os.WriteFile(file, []byte(content), 0o644)
}
