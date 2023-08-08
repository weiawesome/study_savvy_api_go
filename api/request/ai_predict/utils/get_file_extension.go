package utils

import "strings"

func GetFileExtension(filename string) string {
	index := strings.LastIndex(filename, ".")
	if index == -1 {
		return ""
	}
	return strings.ToLower(filename[index:])
}
