package config

import (
	"os"
	"path/filepath"
)

func Executable() string {
	path, _ := os.Executable()
	return filepath.Base(path)
}

func IdLength() int {
	return 6
}
