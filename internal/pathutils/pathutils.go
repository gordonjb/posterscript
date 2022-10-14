package pathutils

import "path/filepath"

func Stem(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}
