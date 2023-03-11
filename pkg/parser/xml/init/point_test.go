package init

import "testing"

func TestPoints_Marshal(t *testing.T) {
	type fields struct {
		Points []Point
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "四人麻雀",
			fields: fields{
				Points: []Point{
					{
						PlayerIndex: 1,
						Point:       25000,
					},
					{
						PlayerIndex: 2,
						Point:       25000,
					},
					{
						PlayerIndex: 3,
						Point:       25000,
					},
					{
						PlayerIndex: 4,
						Point:       25000,
					},
				},
			},
			want: "250,250,250,250",
		},
		{
			name: "三人麻雀",
			fields: fields{
				Points: []Point{
					{
						PlayerIndex: 1,
						Point:       35000,
					},
					{
						PlayerIndex: 2,
						Point:       35000,
					},
					{
						PlayerIndex: 3,
						Point:       35000,
					},
				},
			},
			want: "350,350,350",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Points{
				Points: tt.fields.Points,
			}
			if got := p.Marshal(); got != tt.want {
				t.Errorf("Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoints_Unmarshal(t *testing.T) {
	type fields struct {
		Points []Point
	}
	type args struct {
		point string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Points
		wantErr bool
	}{
		{
			name: "四人麻雀",
			fields: fields{
				Points: []Point{},
			},
			args: args{
				point: "250,250,250,250",
			},
			want: Points{
				Points: []Point{
					{
						PlayerIndex: 1,
						Point:       25000,
					},
					{
						PlayerIndex: 2,
						Point:       25000,
					},
					{
						PlayerIndex: 3,
						Point:       25000,
					},
					{
						PlayerIndex: 4,
						Point:       25000,
					},
				},
			},
		},
		{
			name: "三人麻雀",
			fields: fields{
				Points: []Point{},
			},
			args: args{
				point: "350,350,350",
			},
			want: Points{
				Points: []Point{
					{
						PlayerIndex: 1,
						Point:       35000,
					},
					{
						PlayerIndex: 2,
						Point:       35000,
					},
					{
						PlayerIndex: 3,
						Point:       35000,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Points{
				Points: tt.fields.Points,
			}
			if err := p.Unmarshal(tt.args.point); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
