package operation

import (
	"github.com/google/go-cmp/cmp"
	"github.com/kanade0404/tenhou-log/pkg/parser/xml/hai"
	"log"
	"testing"
)

func TestNewDoraOpen(t *testing.T) {
	hai0, err := hai.NewHai(0, true)
	if err != nil {
		log.Fatalln(err)
	}
	hai3, err := hai.NewHai(3, true)
	if err != nil {
		log.Fatalln(err)
	}
	hai16, err := hai.NewHai(16, true)
	if err != nil {
		log.Fatalln(err)
	}
	hai17, err := hai.NewHai(17, true)
	if err != nil {
		log.Fatalln(err)
	}
	hai35, err := hai.NewHai(35, true)
	if err != nil {
		log.Fatalln(err)
	}
	hai36, err := hai.NewHai(36, true)
	if err != nil {
		log.Fatalln(err)
	}
	hai52, err := hai.NewHai(52, true)
	if err != nil {
		log.Fatalln(err)
	}
	hai53, err := hai.NewHai(53, true)
	if err != nil {
		log.Fatalln(err)
	}
	hai71, err := hai.NewHai(71, true)
	if err != nil {
		log.Fatalln(err)
	}
	hai73, err := hai.NewHai(73, true)
	if err != nil {
		log.Fatalln(err)
	}
	hai88, err := hai.NewHai(88, true)
	if err != nil {
		log.Fatalln(err)
	}
	hai89, err := hai.NewHai(89, true)
	if err != nil {
		log.Fatalln(err)
	}
	hai107, err := hai.NewHai(107, true)
	if err != nil {
		log.Fatalln(err)
	}
	tests := []struct {
		name    string
		haiID   string
		want    hai.IHai
		wantErr bool
	}{
		{
			name:  "一萬",
			haiID: "0",
			want:  hai0,
		},
		{
			name:  "一萬",
			haiID: "3",
			want:  hai3,
		},
		{
			name:  "五萬",
			haiID: "16",
			want:  hai16,
		},
		{
			name:  "五萬",
			haiID: "17",
			want:  hai17,
		},
		{
			name:  "九萬",
			haiID: "35",
			want:  hai35,
		},
		{
			name:  "一筒",
			haiID: "36",
			want:  hai36,
		},
		{
			name:  "五筒",
			haiID: "52",
			want:  hai52,
		},
		{
			name:  "五筒",
			haiID: "53",
			want:  hai53,
		},
		{
			name:  "九筒",
			haiID: "71",
			want:  hai71,
		},
		{
			name:  "一索",
			haiID: "73",
			want:  hai73,
		},
		{
			name:  "五索",
			haiID: "88",
			want:  hai88,
		},
		{
			name:  "五索",
			haiID: "89",
			want:  hai89,
		},
		{
			name:  "九索",
			haiID: "107",
			want:  hai107,
		},
		{
			name:    "空文字でエラー",
			haiID:   "",
			wantErr: true,
		},
		{
			name:    "数字以外でエラー",
			haiID:   "a",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDoraOpen(tt.haiID, true)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("NewDoraOpen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(hai.Hai{})); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
