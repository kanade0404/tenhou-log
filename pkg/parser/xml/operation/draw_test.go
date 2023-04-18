package operation

import (
	"github.com/google/go-cmp/cmp"
	"github.com/kanade0404/tenhou-log/pkg/parser/xml"
	"github.com/kanade0404/tenhou-log/pkg/parser/xml/hai"
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

func TestUnmarshal(t *testing.T) {
	oneCharacter0, err := hai.NewHai(0, true)
	if err != nil {
		t.Error(err)
	}
	//oneCharacter3, err := hai.NewHai(3, true)
	//if err != nil {
	//	t.Error(err)
	//}
	fiveCharacter16, err := hai.NewHai(16, false)
	if err != nil {
		t.Error(err)
	}
	fiveRedCharacter16, err := hai.NewHai(16, true)
	if err != nil {
		t.Error(err)
	}
	//nineCharacter35, err := hai.NewHai(35, true)
	//if err != nil {
	//	t.Error(err)
	//}
	tests := []struct {
		name    string
		e       xml.XmlElement
		isRed   bool
		want    *Draw
		wantErr bool
	}{
		{
			name: "player0 draw 1 character(id:0)",
			e: xml.XmlElement{
				Name: "T0",
			},
			isRed: true,
			want: &Draw{
				DrawType: FirstDrawUser,
				Hai:      oneCharacter0,
				GameInfo: &xml.GameInfo{
					Red: true,
				},
			},
		},
		{
			name: "player1 draw 1 character(id:0)",
			e: xml.XmlElement{
				Name: "U0",
			},
			isRed: true,
			want: &Draw{
				DrawType: SecondDrawUser,
				Hai:      oneCharacter0,
				GameInfo: &xml.GameInfo{
					Red: true,
				},
			},
		},
		{
			name: "player2 draw 1 character(id:0)",
			e: xml.XmlElement{
				Name: "V0",
			},
			isRed: true,
			want: &Draw{
				DrawType: ThirdDrawUser,
				Hai:      oneCharacter0,
				GameInfo: &xml.GameInfo{
					Red: true,
				},
			},
		},
		{
			name: "player3 draw 1 character(id:0)",
			e: xml.XmlElement{
				Name: "W0",
			},
			isRed: true,
			want: &Draw{
				DrawType: FourthDrawUser,
				Hai:      oneCharacter0,
				GameInfo: &xml.GameInfo{
					Red: true,
				},
			},
		},
		{
			name: "player0 draw 5 character(id:16)",
			e: xml.XmlElement{
				Name: "T16",
			},
			isRed: false,
			want: &Draw{
				DrawType: FirstDrawUser,
				Hai:      fiveCharacter16,
				GameInfo: &xml.GameInfo{
					Red: false,
				},
			},
		},
		{
			name: "player1 draw 5 character(id:16)",
			e: xml.XmlElement{
				Name: "U16",
			},
			isRed: false,
			want: &Draw{
				DrawType: SecondDrawUser,
				Hai:      fiveCharacter16,
				GameInfo: &xml.GameInfo{
					Red: false,
				},
			},
		},
		{
			name: "player2 draw 5 character(id:16)",
			e: xml.XmlElement{
				Name: "V16",
			},
			isRed: false,
			want: &Draw{
				DrawType: ThirdDrawUser,
				Hai:      fiveCharacter16,
				GameInfo: &xml.GameInfo{
					Red: false,
				},
			},
		},
		{
			name: "player3 draw 5 character(id:16)",
			e: xml.XmlElement{
				Name: "W16",
			},
			isRed: false,
			want: &Draw{
				DrawType: FourthDrawUser,
				Hai:      fiveCharacter16,
				GameInfo: &xml.GameInfo{
					Red: false,
				},
			},
		},
		{
			name: "player0 draw red5 character(id:16)",
			e: xml.XmlElement{
				Name: "T16",
			},
			isRed: true,
			want: &Draw{
				DrawType: FirstDrawUser,
				Hai:      fiveRedCharacter16,
				GameInfo: &xml.GameInfo{
					Red: true,
				},
			},
		},
		{
			name: "player1 draw red5 character(id:16)",
			e: xml.XmlElement{
				Name: "U16",
			},
			isRed: true,
			want: &Draw{
				DrawType: SecondDrawUser,
				Hai:      fiveRedCharacter16,
				GameInfo: &xml.GameInfo{
					Red: true,
				},
			},
		},
		{
			name: "player2 draw red5 character(id:16)",
			e: xml.XmlElement{
				Name: "V16",
			},
			isRed: true,
			want: &Draw{
				DrawType: ThirdDrawUser,
				Hai:      fiveRedCharacter16,
				GameInfo: &xml.GameInfo{
					Red: true,
				},
			},
		},
		{
			name: "player3 draw red5 character(id:16)",
			e: xml.XmlElement{
				Name: "W16",
			},
			isRed: true,
			want: &Draw{
				DrawType: FourthDrawUser,
				Hai:      fiveRedCharacter16,
				GameInfo: &xml.GameInfo{
					Red: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Unmarshal(tt.e, &xml.GameInfo{Red: tt.isRed})
			if err != nil {
				if tt.wantErr {
					return
				} else {
					t.Errorf("NewDraw() error=%v, wantErr %t", err, tt.isRed)
				}
			}
			if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(hai.Hai{})); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
