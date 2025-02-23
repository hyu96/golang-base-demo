package file

import (
	"path/filepath"
	"strings"
)

func ValidateCsvExt(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	if ext == ".csv" {
		return true
	}
	return false
}
