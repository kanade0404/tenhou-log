package operation

import (
	"github.com/kanade0404/tenhou-log/serdes/xml"
	"github.com/kanade0404/tenhou-log/serdes/xml/hai"
	"reflect"
	"testing"
)

func TestDrawType_PlayerIndex(t *testing.T) {
	tests := []struct {
		name string
		t    DrawType
		want uint
	}{
		{
			name: "対局開始時東家=0",
			t:    FirstDrawUser,
			want: 0,
		},
		{
			name: "対局開始時南家=1",
			t:    SecondDrawUser,
			want: 1,
		},
		{
			name: "対局開始時西家=2",
			t:    ThirdDrawUser,
			want: 2,
		},
		{
			name: "対局開始時北家=3",
			t:    FourthDrawUser,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.PlayerIndex(); got != tt.want {
				t.Errorf("PlayerIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDrawType(t *testing.T) {
	type args struct {
		drawType string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "'T'はDrawType",
			args: args{
				drawType: "T",
			},
			want: true,
		},
		{
			name: "'U'はDrawType",
			args: args{
				drawType: "U",
			},
			want: true,
		},
		{
			name: "'V'はDrawType",
			args: args{
				drawType: "V",
			},
			want: true,
		},
		{
			name: "'W'はDrawType",
			args: args{
				drawType: "W",
			},
			want: true,
		},
		{
			name: "'X'はDrawTypeではない",
			args: args{
				drawType: "X",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDrawType(tt.args.drawType); got != tt.want {
				t.Errorf("IsDrawType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDraw(t *testing.T) {
	type args struct {
		drawType string
		haiID    int
		gameInfo *xml.GameInfo
	}
	red5Character, _ := hai.NewHai(16, true)
	nonRed5Character, _ := hai.NewHai(16, false)
	tests := []struct {
		name    string
		args    args
		want    *Draw
		wantErr bool
	}{
		{
			name: "赤ありで対局開始時東家のツモは赤五萬",
			args: args{
				drawType: "T",
				haiID:    16,
				gameInfo: &xml.GameInfo{
					Red: true,
				},
			},
			want: &Draw{
				DrawType: FirstDrawUser,
				Hai:      red5Character,
				GameInfo: &xml.GameInfo{
					Red: true,
				},
			},
		},
		{
			name: "赤なしで対局開始時東家のツモは五萬",
			args: args{
				drawType: "T",
				haiID:    16,
				gameInfo: &xml.GameInfo{
					Red: false,
				},
			},
			want: &Draw{
				DrawType: FirstDrawUser,
				Hai:      nonRed5Character,
				GameInfo: &xml.GameInfo{
					Red: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDraw(tt.args.drawType, tt.args.haiID, tt.args.gameInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDraw() got = %v, want %v", got, tt.want)
			}
		})
	}
}
