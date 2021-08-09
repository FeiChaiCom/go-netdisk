package misc

import (
	"os"
	"path/filepath"
	"strings"
)

//FileExists check file exists
func FileExists(filename string) (bool, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

//FileName return filename without extension
func FileName(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

//SplitDirFile return dir and file of the filepath
func SplitDirFile(filePath string) (string, string) {
	dir := filepath.Dir(filePath)
	name := filepath.Base(filePath)
	return dir, name
}
