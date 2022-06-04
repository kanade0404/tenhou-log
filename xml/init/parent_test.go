package init

import "testing"

func TestParentPlayer_Marshal(t *testing.T) {
	type fields struct {
		Index uint
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "正常系(0)",
			fields: fields{
				Index: 1,
			},
			want: "0",
		},
		{
			name: "正常系(4)",
			fields: fields{
				Index: 4,
			},
			want: "3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ParentPlayer{
				Index: tt.fields.Index,
			}
			if got := p.Marshal(); got != tt.want {
				t.Errorf("Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParentPlayer_Unmarshal(t *testing.T) {
	type fields struct {
		Index uint
	}
	type args struct {
		index string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ParentPlayer
		wantErr bool
	}{
		{
			name: "正常系(0)",
			fields: fields{
				Index: 0,
			},
			args: args{
				index: "0",
			},
			want: ParentPlayer{
				Index: 1,
			},
		},
		{
			name: "正常系(4)",
			fields: fields{
				Index: 0,
			},
			args: args{
				index: "3",
			},
			want: ParentPlayer{
				Index: 4,
			},
		},
		{
			name: "正常系(3)",
			fields: fields{
				Index: 0,
			},
			args: args{
				index: "3",
			},
			want: ParentPlayer{
				Index: 4,
			},
		},
		{
			name: "異常系(5)",
			fields: fields{
				Index: 0,
			},
			args: args{
				index: "4",
			},
			wantErr: true,
		},
		{
			name: "異常系(非数値)",
			fields: fields{
				Index: 0,
			},
			args: args{
				index: "a",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ParentPlayer{
				Index: tt.fields.Index,
			}
			if err := p.Unmarshal(tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
