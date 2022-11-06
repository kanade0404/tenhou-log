package hai

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestHai_Name(t *testing.T) {
	type fields struct {
		ID   uint
		Num  uint
		Type HaiType
	}
	tests := []struct {
		name   string
		fields fields
		isRed  bool
		want   string
	}{
		{
			name: "一萬赤なし",
			fields: fields{
				ID:   0,
				Num:  1,
				Type: CharactersType,
			},
			want: "一萬",
		},
		{
			name: "一萬赤なし",
			fields: fields{
				ID:   3,
				Num:  1,
				Type: CharactersType,
			},
			want: "一萬",
		},
		{
			name: "五萬赤なし",
			fields: fields{
				ID:   16,
				Num:  5,
				Type: CharactersType,
			},
			want: "五萬",
		},
		{
			name: "五萬赤あり（赤五萬）",
			fields: fields{
				ID:   16,
				Num:  5,
				Type: CharactersType,
			},
			isRed: true,
			want:  "赤五萬",
		},
		{
			name: "五萬赤あり（黒五萬）",
			fields: fields{
				ID:   17,
				Num:  5,
				Type: CharactersType,
			},
			isRed: true,
			want:  "五萬",
		},
		{
			name: "九萬赤なし",
			fields: fields{
				ID:   35,
				Num:  9,
				Type: CharactersType,
			},
			want: "九萬",
		},
		{
			name: "一筒赤なし",
			fields: fields{
				ID:   36,
				Num:  1,
				Type: CirclesType,
			},
			want: "一筒",
		},
		{
			name: "五筒赤なし",
			fields: fields{
				ID:   52,
				Num:  5,
				Type: CirclesType,
			},
			want: "五筒",
		},
		{
			name: "五筒赤あり（赤五筒）",
			fields: fields{
				ID:   52,
				Num:  5,
				Type: CirclesType,
			},
			isRed: true,
			want:  "赤五筒",
		},
		{
			name: "五筒赤あり（黒五筒）",
			fields: fields{
				ID:   53,
				Num:  5,
				Type: CirclesType,
			},
			isRed: true,
			want:  "五筒",
		},
		{
			name: "九筒赤なし",
			fields: fields{
				ID:   71,
				Num:  9,
				Type: CirclesType,
			},
			want: "九筒",
		},
		{
			name: "一索赤なし",
			fields: fields{
				ID:   73,
				Num:  1,
				Type: BamboosType,
			},
			want: "一索",
		},
		{
			name: "五索赤なし",
			fields: fields{
				ID:   88,
				Num:  5,
				Type: BamboosType,
			},
			want: "五索",
		},
		{
			name: "五索赤あり（赤五索）",
			fields: fields{
				ID:   88,
				Num:  5,
				Type: BamboosType,
			},
			isRed: true,
			want:  "赤五索",
		},
		{
			name: "五索赤あり（黒五索）",
			fields: fields{
				ID:   89,
				Num:  5,
				Type: BamboosType,
			},
			isRed: true,
			want:  "五索",
		},
		{
			name: "九索赤なし",
			fields: fields{
				ID:   107,
				Num:  9,
				Type: BamboosType,
			},
			isRed: true,
			want:  "九索",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, _ := NewHai(tt.fields.ID, tt.isRed)
			if got := h.Name(); got != tt.want {
				t.Errorf("File() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHai(t *testing.T) {
	type args struct {
		id    uint
		isRed bool
	}
	tests := []struct {
		name    string
		args    args
		want    IHai
		wantErr bool
	}{
		{
			name: "一萬(0)赤あり",
			args: args{
				id:    0,
				isRed: true,
			},
			want: &Characters{
				Hai: &Hai{
					id:        0,
					num:       1,
					haiType:   CharactersType,
					isRedRule: true,
				},
			},
		},
		{
			name: "一萬(3)赤なし",
			args: args{
				id:    3,
				isRed: false,
			},
			want: &Characters{
				Hai: &Hai{
					id:        3,
					num:       1,
					haiType:   CharactersType,
					isRedRule: false,
				},
			},
		},
		{
			name: "九萬(32)赤あり",
			args: args{
				id:    32,
				isRed: true,
			},
			want: &Characters{
				Hai: &Hai{
					id:        32,
					num:       9,
					haiType:   CharactersType,
					isRedRule: true,
				},
			},
		},
		{
			name: "九萬(35)赤なし",
			args: args{
				id:    35,
				isRed: false,
			},
			want: &Characters{
				Hai: &Hai{
					id:        35,
					num:       9,
					haiType:   CharactersType,
					isRedRule: false,
				},
			},
		},
		{
			name: "一筒(36)赤あり",
			args: args{
				id:    36,
				isRed: true,
			},
			want: &Circles{
				Hai: &Hai{
					id:        36,
					num:       1,
					haiType:   CirclesType,
					isRedRule: true,
				},
			},
		},
		{
			name: "一筒(37)赤なし",
			args: args{
				id:    37,
				isRed: false,
			},
			want: &Circles{
				Hai: &Hai{
					id:        37,
					num:       1,
					haiType:   CirclesType,
					isRedRule: false,
				},
			},
		},
		{
			name: "九筒(70)赤あり",
			args: args{
				id:    70,
				isRed: true,
			},
			want: &Circles{
				Hai: &Hai{
					id:        70,
					num:       9,
					haiType:   CirclesType,
					isRedRule: true,
				},
			},
		},
		{
			name: "九筒(71)赤なし",
			args: args{
				id:    71,
				isRed: false,
			},
			want: &Circles{
				Hai: &Hai{
					id:        71,
					num:       9,
					haiType:   CirclesType,
					isRedRule: false,
				},
			},
		},
		{
			name: "一索(72)赤あり",
			args: args{
				id:    72,
				isRed: true,
			},
			want: &Bamboos{
				Hai: &Hai{
					id:        72,
					num:       1,
					haiType:   BamboosType,
					isRedRule: true,
				},
			},
		},
		{
			name: "一索(73)赤なし",
			args: args{
				id:    73,
				isRed: false,
			},
			want: &Bamboos{
				Hai: &Hai{
					id:        73,
					num:       1,
					haiType:   BamboosType,
					isRedRule: false,
				},
			},
		},
		{
			name: "九索(106)赤あり",
			args: args{
				id:    106,
				isRed: true,
			},
			want: &Bamboos{
				Hai: &Hai{
					id:        106,
					num:       9,
					haiType:   BamboosType,
					isRedRule: true,
				},
			},
		},
		{
			name: "九索(107)赤なし",
			args: args{
				id:    107,
				isRed: false,
			},
			want: &Bamboos{
				Hai: &Hai{
					id:        107,
					num:       9,
					haiType:   BamboosType,
					isRedRule: false,
				},
			},
		},
		{
			name: "東(108)",
			args: args{
				id:    108,
				isRed: false,
			},
			want: &Honours{
				Hai: Hai{
					id:        108,
					num:       1,
					haiType:   HonorsType,
					isRedRule: false,
				},
			},
		},
		{
			name: "東(111)",
			args: args{
				id:    111,
				isRed: false,
			},
			want: &Honours{
				Hai: Hai{
					id:        111,
					num:       1,
					haiType:   HonorsType,
					isRedRule: false,
				},
			},
		}, {
			name: "南(112)",
			args: args{
				id:    112,
				isRed: false,
			},
			want: &Honours{
				Hai: Hai{
					id:        112,
					num:       2,
					haiType:   HonorsType,
					isRedRule: false,
				},
			},
		},
		{
			name: "南(115)",
			args: args{
				id:    115,
				isRed: false,
			},
			want: &Honours{
				Hai: Hai{
					id:        115,
					num:       2,
					haiType:   HonorsType,
					isRedRule: false,
				},
			},
		},
		{
			name: "西(116)",
			args: args{
				id:    116,
				isRed: false,
			},
			want: &Honours{
				Hai: Hai{
					id:        116,
					num:       3,
					haiType:   HonorsType,
					isRedRule: false,
				},
			},
		},
		{
			name: "西(119)",
			args: args{
				id:    119,
				isRed: false,
			},
			want: &Honours{
				Hai: Hai{
					id:        119,
					num:       3,
					haiType:   HonorsType,
					isRedRule: false,
				},
			},
		},
		{
			name: "北(120)",
			args: args{
				id:    120,
				isRed: false,
			},
			want: &Honours{
				Hai: Hai{
					id:        120,
					num:       4,
					haiType:   HonorsType,
					isRedRule: false,
				},
			},
		},
		{
			name: "北(123)",
			args: args{
				id:    123,
				isRed: false,
			},
			want: &Honours{
				Hai: Hai{
					id:        123,
					num:       4,
					haiType:   HonorsType,
					isRedRule: false,
				},
			},
		},
		{
			name: "白(124)",
			args: args{
				id:    124,
				isRed: false,
			},
			want: &Honours{
				Hai: Hai{
					id:        124,
					num:       5,
					haiType:   HonorsType,
					isRedRule: false,
				},
			},
		},
		{
			name: "白(127)",
			args: args{
				id:    127,
				isRed: false,
			},
			want: &Honours{
				Hai: Hai{
					id:        127,
					num:       5,
					haiType:   HonorsType,
					isRedRule: false,
				},
			},
		},
		{
			name: "發(128)",
			args: args{
				id:    128,
				isRed: false,
			},
			want: &Honours{
				Hai: Hai{
					id:        128,
					num:       6,
					haiType:   HonorsType,
					isRedRule: false,
				},
			},
		},
		{
			name: "發(131)",
			args: args{
				id:    131,
				isRed: false,
			},
			want: &Honours{
				Hai: Hai{
					id:        131,
					num:       6,
					haiType:   HonorsType,
					isRedRule: false,
				},
			},
		},
		{
			name: "中(132)",
			args: args{
				id:    132,
				isRed: false,
			},
			want: &Honours{
				Hai: Hai{
					id:        132,
					num:       7,
					haiType:   HonorsType,
					isRedRule: false,
				},
			},
		},
		{
			name: "中(135)",
			args: args{
				id:    135,
				isRed: false,
			},
			want: &Honours{
				Hai: Hai{
					id:        135,
					num:       7,
					haiType:   HonorsType,
					isRedRule: false,
				},
			},
		},
		{
			name: "字牌の上限135を上回るのでエラー",
			args: args{
				id:    136,
				isRed: false,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewHai(tt.args.id, tt.args.isRed)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHai() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(fmt.Sprintf("%+v", got), fmt.Sprintf("%+v", tt.want)); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}

func Test_IsRed(t *testing.T) {
	tests := []struct {
		name      string
		id        uint
		isRedRule bool
		want      bool
		wantErr   bool
	}{
		{
			name:      "赤なし一萬(0)",
			id:        0,
			isRedRule: false,
			want:      false,
		},
		{
			name:      "赤あり一萬(0)",
			id:        0,
			isRedRule: true,
			want:      false,
		},
		{
			name:      "赤なし五萬(15)",
			id:        15,
			isRedRule: false,
			want:      false,
		},
		{
			name:      "赤あり五萬(15)",
			id:        15,
			isRedRule: true,
			want:      false,
		},
		{
			name:      "赤なし五萬(16)",
			id:        16,
			isRedRule: false,
			want:      false,
		},
		{
			name:      "赤あり五萬(16)",
			id:        16,
			isRedRule: true,
			want:      true,
		},
		{
			name:      "赤なし五筒(51)",
			id:        51,
			isRedRule: false,
			want:      false,
		},
		{
			name:      "赤あり五筒(51)",
			id:        51,
			isRedRule: true,
			want:      false,
		},
		{
			name:      "赤なし五筒(52)",
			id:        52,
			isRedRule: false,
			want:      false,
		},
		{
			name:      "赤あり五筒(52)",
			id:        52,
			isRedRule: true,
			want:      true,
		},
		{
			name:      "赤なし五索(87)",
			id:        87,
			isRedRule: false,
			want:      false,
		},
		{
			name:      "赤あり五索(87)",
			id:        87,
			isRedRule: true,
			want:      false,
		},
		{
			name:      "赤なし五索(88)",
			id:        88,
			isRedRule: false,
			want:      false,
		},
		{
			name:      "赤あり五索(88)",
			id:        88,
			isRedRule: true,
			want:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, err := NewHai(tt.id, tt.isRedRule)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsRed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got := h.IsRed(); got != tt.want {
				t.Errorf("IsRed() = %v, want %v", got, tt.want)
			}
		})
	}
}
