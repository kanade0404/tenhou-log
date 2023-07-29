package internal

import (
	"context"
	"regexp"
	"strings"

	tracer2 "github.com/kanade0404/tenhou-log/pkg/driver/tracer"
)

var tracer = tracer2.NewTracer("services/scraper/internal/unmarshal_file_list_text")

func UnmarshalFileListText(ctx context.Context, text string) []string {
	_, span := tracer.Start(ctx, "UnmarshalFileListText")
	defer span.End()
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
