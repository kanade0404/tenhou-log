package hai

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBamboos_Name(t *testing.T) {
	type fields struct {
		Suits Suits
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "1索赤なし",
			fields: fields{
				Suits{
					Hai{
						ID:   73,
						Num:  1,
						Type: BamboosType,
					},
					false,
				},
			},
			want: "一索",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s ID:%d Num:%d Type:%s", tt.name, tt.fields.Suits.ID, tt.fields.Suits.Num, tt.fields.Suits.Type), func(t *testing.T) {
			b := Bamboos{
				Suits: tt.fields.Suits,
			}
			if got := b.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharacters_Name(t *testing.T) {
	type fields struct {
		Suits Suits
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Characters{
				Suits: tt.fields.Suits,
			}
			if got := c.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircles_Name(t *testing.T) {
	type fields struct {
		Suits Suits
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circles{
				Suits: tt.fields.Suits,
			}
			if got := c.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHai_Name(t *testing.T) {
	type fields struct {
		ID   uint
		Num  uint
		Type HaiType
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hai{
				ID:   tt.fields.ID,
				Num:  tt.fields.Num,
				Type: tt.fields.Type,
			}
			if got := h.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHonours_Name(t *testing.T) {
	type fields struct {
		Hai Hai
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Honours{
				Hai: tt.fields.Hai,
			}
			if got := h.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHai(t *testing.T) {
	type args struct {
		num   uint
		isRed bool
	}
	tests := []struct {
		name    string
		args    args
		want    IHai
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewHai(tt.args.num, tt.args.isRed)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHai() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHai() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuits_Name(t *testing.T) {
	type fields struct {
		Hai   Hai
		IsRed bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		//
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Suits{
				Hai:   tt.fields.Hai,
				IsRed: tt.fields.IsRed,
			}
			if got := s.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newBamboos(t *testing.T) {
	type args struct {
		id    uint
		isRed bool
	}
	tests := []struct {
		name    string
		args    args
		want    IBamboos
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newBamboos(tt.args.id, tt.args.isRed)
			if (err != nil) != tt.wantErr {
				t.Errorf("newBamboos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newBamboos() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newCharacters(t *testing.T) {
	type args struct {
		id    uint
		isRed bool
	}
	tests := []struct {
		name    string
		args    args
		want    ICharacters
		wantErr bool
	}{
		{
			name: "一萬(0)赤あり",
			args: args{
				id:    0,
				isRed: true,
			},
			want: Characters{
				Suits: Suits{
					Hai: Hai{
						ID:   0,
						Num:  1,
						Type: CharactersType,
					},
					IsRed: true,
				},
			},
		},
		{
			name: "一萬(3)赤なし",
			args: args{
				id:    3,
				isRed: false,
			},
			want: Characters{
				Suits: Suits{
					Hai: Hai{
						ID:   3,
						Num:  1,
						Type: CharactersType,
					},
					IsRed: false,
				},
			},
		},
		{
			name: "九萬(32)赤あり",
			args: args{
				id:    32,
				isRed: true,
			},
			want: Characters{
				Suits: Suits{
					Hai: Hai{
						ID:   32,
						Num:  9,
						Type: CharactersType,
					},
					IsRed: true,
				},
			},
		},
		{
			name: "九萬(35)赤なし",
			args: args{
				id:    35,
				isRed: false,
			},
			want: Characters{
				Suits: Suits{
					Hai: Hai{
						ID:   35,
						Num:  9,
						Type: CharactersType,
					},
					IsRed: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newCharacters(tt.args.id, tt.args.isRed)
			if (err != nil) != tt.wantErr {
				t.Errorf("newCharacters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCharacters() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newCircles(t *testing.T) {
	type args struct {
		id    uint
		isRed bool
	}
	tests := []struct {
		name    string
		args    args
		want    ICircles
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newCircles(tt.args.id, tt.args.isRed)
			if (err != nil) != tt.wantErr {
				t.Errorf("newCircles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCircles() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newHonours(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		args    args
		want    IHonours
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newHonours(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("newHonours() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newHonours() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newSuits(t *testing.T) {
	type args struct {
		id    uint
		isRed bool
	}
	tests := []struct {
		name    string
		args    args
		want    ISuits
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newSuits(tt.args.id, tt.args.isRed)
			if (err != nil) != tt.wantErr {
				t.Errorf("newSuits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSuits() got = %v, want %v", got, tt.want)
			}
		})
	}
}
