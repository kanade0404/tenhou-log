package init

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/kanade0404/tenhou-log/serdes/xml"
	"github.com/kanade0404/tenhou-log/serdes/xml/hai"
	"sort"
	"testing"
)

func TestHandHais_Marshal(t *testing.T) {
	var (
		test1Hais []hai.IHai
		test2Hais []hai.IHai
		test3Hais []hai.IHai
	)
	type fields struct {
		Hands []hai.IHai
		_     struct{}
	}
	for i := 0; i < 13; i++ {
		h, err := hai.NewHai(i, false)
		if err != nil {
			t.Fatal(err)
		}
		test1Hais = append(test1Hais, h)
	}
	for i := 52; i < 65; i++ {
		h, err := hai.NewHai(i, true)
		if err != nil {
			t.Fatal(err)
		}
		test2Hais = append(test2Hais, h)
	}
	for i := 72; i < 85; i++ {
		h, err := hai.NewHai(i, true)
		if err != nil {
			t.Fatal(err)
		}
		test3Hais = append(test3Hais, h)
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test1",
			fields: fields{
				Hands: test1Hais,
			},
			want: "0,1,2,3,4,5,6,7,8,9,10,11,12",
		},
		{
			name: "test2",
			fields: fields{
				Hands: test2Hais,
			},
			want: "52,53,54,55,56,57,58,59,60,61,62,63,64",
		},
		{
			name: "test3",
			fields: fields{
				Hands: test3Hais,
			},
			want: "72,73,74,75,76,77,78,79,80,81,82,83,84",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HandHais{
				Hands: tt.fields.Hands,
			}
			if got := h.Marshal(); got != tt.want {
				t.Errorf("Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandHais_Unmarshal(t *testing.T) {
	var (
		test1Hais []hai.IHai
		test2Hais []hai.IHai
		test3Hais []hai.IHai
	)
	for i := 0; i < 13; i++ {
		h, err := hai.NewHai(i, false)
		if err != nil {
			t.Fatal(err)
		}
		test1Hais = append(test1Hais, h)
	}
	for i := 52; i < 65; i++ {
		h, err := hai.NewHai(i, true)
		if err != nil {
			t.Fatal(err)
		}
		test2Hais = append(test2Hais, h)
	}
	for i := 72; i < 85; i++ {
		h, err := hai.NewHai(i, true)
		if err != nil {
			t.Fatal(err)
		}
		test3Hais = append(test3Hais, h)
	}
	type fields struct {
		Hands    []hai.IHai
		gameInfo *xml.GameInfo
		_        struct{}
	}
	type args struct {
		hais string
		_    struct{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    HandHais
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				Hands:    []hai.IHai{},
				gameInfo: &xml.GameInfo{},
			},
			args: args{
				hais: "0,1,2,3,4,5,6,7,8,9,10,11,12",
			},
			want: HandHais{
				Hands: test1Hais,
			},
		},
		{
			name: "test2",
			fields: fields{
				Hands: []hai.IHai{},
				gameInfo: &xml.GameInfo{
					Red: true,
				},
			},
			args: args{
				hais: "52,53,54,55,56,57,58,59,60,61,62,63,64",
			},
			want: HandHais{
				Hands: test2Hais,
			},
		},
		{
			name: "test3",
			fields: fields{
				Hands: []hai.IHai{},
				gameInfo: &xml.GameInfo{
					Red: true,
				},
			},
			args: args{
				hais: "72,73,74,75,76,77,78,79,80,81,82,83,84",
			},
			want: HandHais{
				Hands: test3Hais,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandHais{
				Hands:    tt.fields.Hands,
				gameInfo: tt.fields.gameInfo,
			}
			if err := h.Unmarshal(tt.args.hais); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			sort.Slice(h.Hands, func(i, j int) bool {
				return h.Hands[i].ID() < h.Hands[j].ID()
			})
			sort.Slice(tt.want.Hands, func(i, j int) bool {
				return tt.want.Hands[i].ID() < tt.want.Hands[j].ID()
			})
			if diff := cmp.Diff(fmt.Sprintf("%v", h.Hands), fmt.Sprintf("%v", tt.want.Hands)); diff != "" {
				t.Errorf("(-got+want\n%v)", diff)
			}
		})
	}
}
