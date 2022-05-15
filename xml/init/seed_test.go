package init

import "testing"

func TestHonba_String(t *testing.T) {
	tests := []struct {
		name string
		h    Honba
		want string
	}{
		{
			name: "0本場",
			h:    Honba(0),
			want: "0本場",
		},
		{
			name: "1本場",
			h:    Honba(1),
			want: "1本場",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDiceValue(t *testing.T) {
	type args struct {
		diceValue uint
	}
	tests := []struct {
		name    string
		args    args
		want    DiceValue
		wantErr bool
	}{
		{
			name: "正常値:0",
			args: args{
				diceValue: 0,
			},
			want:    DiceValue(1),
			wantErr: false,
		},
		{
			name: "正常値:5",
			args: args{
				diceValue: 5,
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "異常値:6",
			args: args{
				diceValue: 6,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDiceValue(tt.args.diceValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDiceValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewDiceValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRound(t *testing.T) {
	type args struct {
		round uint
	}
	tests := []struct {
		name string
		args args
		want Round
	}{
		{
			name: "東場",
			args: args{
				round: 0,
			},
			want: Round("東"),
		},
		{
			name: "南場",
			args: args{
				round: 1,
			},
			want: Round("南"),
		},
		{
			name: "西場",
			args: args{
				round: 2,
			},
			want: Round("西"),
		},
		{
			name: "北場",
			args: args{
				round: 3,
			},
			want: Round("北"),
		},
		{
			name: "不明",
			args: args{
				round: 4,
			},
			want: Round("UNKNOWN"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRound(tt.args.round); got != tt.want {
				t.Errorf("NewRound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeed_Marshal(t *testing.T) {
	type fields struct {
		Round       Round
		Hand        Hand
		Honba       Honba
		RiichPoints RiichPoints
		DiceValue1  DiceValue
		DiceValue2  DiceValue
	}
	tests := []struct {
		name     string
		fields   fields
		wantSeed string
		wantErr  bool
	}{
		{
			name: "東1局0本場ダイスは1と3",
			fields: fields{
				Round:       East,
				Hand:        First,
				RiichPoints: 0,
				DiceValue1:  1,
				DiceValue2:  3,
			},
			wantSeed: "0,0,0,0,0,2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Seed{
				Round:       tt.fields.Round,
				Hand:        tt.fields.Hand,
				Honba:       tt.fields.Honba,
				RiichPoints: tt.fields.RiichPoints,
				DiceValue1:  tt.fields.DiceValue1,
				DiceValue2:  tt.fields.DiceValue2,
			}
			gotSeed, err := s.Marshal()
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSeed != tt.wantSeed {
				t.Errorf("Marshal() gotSeed = %v, want %v", gotSeed, tt.wantSeed)
			}
		})
	}
}

func TestSeed_Unmarshal(t *testing.T) {
	type fields struct {
		Round       Round
		Hand        Hand
		Honba       Honba
		RiichPoints RiichPoints
		DiceValue1  DiceValue
		DiceValue2  DiceValue
	}
	type args struct {
		seed string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Seed{
				Round:       tt.fields.Round,
				Hand:        tt.fields.Hand,
				Honba:       tt.fields.Honba,
				RiichPoints: tt.fields.RiichPoints,
				DiceValue1:  tt.fields.DiceValue1,
				DiceValue2:  tt.fields.DiceValue2,
			}
			if err := s.Unmarshal(tt.args.seed); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnmarshalRound(t *testing.T) {
	type args struct {
		round Round
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalRound(tt.args.round)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalRound() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UnmarshalRound() got = %v, want %v", got, tt.want)
			}
		})
	}
}
