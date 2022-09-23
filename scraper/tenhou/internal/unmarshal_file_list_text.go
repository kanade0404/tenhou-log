package internal

import (
	"regexp"
	"strings"
)

func UnmarshalFileListText(text string) []string {
	re := regexp.MustCompile(`list|[()\[\];\\R]`)
	arr := strings.Split(strings.ReplaceAll(strings.ReplaceAll(re.ReplaceAllString(text, ""), "file", `"file"`), "size", `"size"`), "\n")
	var result []string
	for _, a := range arr {
		if a != "" {
			if strings.HasSuffix(a, ",") {
				result = append(result, strings.TrimSuffix(a, ","))
			} else {
				result = append(result, a)
			}
		}
	}
	return result
}
