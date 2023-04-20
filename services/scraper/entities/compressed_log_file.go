package entities

import (
	"context"
	"encoding/json"
	"go.opentelemetry.io/otel"
	"strings"
)

type CompressedLogFile struct {
	File string `json:"file"`
	Size uint   `json:"size"`
}

var tracer = otel.GetTracerProvider().Tracer("github.com/kanade0404/tenhou-log/services/scraper/entities/compressed_log_file")

func Unmarshal(ctx context.Context, logText string) (*CompressedLogFile, error) {
	_, span := tracer.Start(ctx, "Unmarshal")
	defer span.End()
	var l CompressedLogFile
	t := strings.ReplaceAll(strings.ReplaceAll(strings.Replace(logText, "},", "}", 1), "\n", ""), "'", `"`)
	err := json.Unmarshal([]byte(t), &l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}
