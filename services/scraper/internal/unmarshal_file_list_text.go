package internal

import (
	"context"
	"go.opentelemetry.io/otel"
	"regexp"
	"strings"
)

var tracer = otel.GetTracerProvider().Tracer("github.com/kanade0404/tenhou-log/services/scraper/internal/unmarshal_file_list_text")

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
