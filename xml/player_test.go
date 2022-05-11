package xml

import (
	"github.com/google/go-cmp/cmp"
	"sort"
	"testing"
)

type TestPlayer struct {
	Name
	Rate
	Dan
	Sex
}
type Name struct {
	original PlayerName
	encoded  string
}

var (
	n0 = TestPlayer{
		Name: Name{
			original: "リンネン",
			encoded:  "%E3%83%AA%E3%83%B3%E3%83%8D%E3%83%B3",
		},
		Rate: 2054.10,
		Dan:  16,
		Sex:  "F",
	}
	n1 = TestPlayer{
		Name: Name{
			original: "上海ニック",
			encoded:  "%E4%B8%8A%E6%B5%B7%E3%83%8B%E3%83%83%E3%82%AF",
		},
		Rate: 2159.83,
		Dan:  16,
		Sex:  "M",
	}
	n2 = TestPlayer{
		Name: Name{
			original: "ekarute",
			encoded:  "%65%6B%61%72%75%74%65",
		},
		Rate: 2159.83,
		Dan:  16,
		Sex:  "M",
	}
	n3 = TestPlayer{
		Name: Name{
			original: "ノオチャベス",
			encoded:  "%E3%83%8E%E3%82%AA%E3%83%81%E3%83%A3%E3%83%99%E3%82%B9",
		},
		Rate: 2204.17,
		Dan:  17,
		Sex:  "M",
	}
)

func TestDecodePlayerName(t *testing.T) {
	type args struct {
		encoded string
	}
	tests := []struct {
		name    string
		args    args
		want    PlayerName
		wantErr bool
	}{
		{
			name:    string(n0.original),
			args:    args{n0.encoded},
			want:    n0.original,
			wantErr: false,
		},
		{
			name:    string(n1.original),
			args:    args{n1.encoded},
			want:    n1.original,
			wantErr: false,
		},
		{
			name:    string(n2.original),
			args:    args{n2.encoded},
			want:    n2.original,
			wantErr: false,
		},
		{
			name:    string(n3.original),
			args:    args{n3.encoded},
			want:    n3.original,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodePlayerName(tt.args.encoded)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodePlayerName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodePlayerName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPlayerIndex(t *testing.T) {
	type args struct {
		idx int
	}
	tests := []struct {
		name    string
		args    args
		want    PlayerIndex
		wantErr bool
	}{
		{
			name:    "正常値:0",
			args:    args{idx: 0},
			want:    0,
			wantErr: false,
		},
		{
			name:    "正常値:3",
			args:    args{idx: 3},
			want:    3,
			wantErr: false,
		},
		{
			name:    "異常値:-1",
			args:    args{idx: -1},
			wantErr: true,
		},
		{
			name:    "異常値:4",
			args:    args{idx: 4},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlayerIndex(tt.args.idx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlayerIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewPlayerIndex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPlayerIndexFromString(t *testing.T) {
	type args struct {
		pi string
	}
	tests := []struct {
		name    string
		args    args
		want    PlayerIndex
		wantErr bool
	}{
		{
			name:    "正常値:0",
			args:    args{pi: "0"},
			want:    0,
			wantErr: false,
		},
		{
			name:    "正常値:3",
			args:    args{pi: "3"},
			want:    3,
			wantErr: false,
		},
		{
			name:    "異常値:-1",
			args:    args{pi: "-1"},
			wantErr: true,
		},
		{
			name:    "異常値:4",
			args:    args{pi: "4"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlayerIndexFromString(tt.args.pi)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlayerIndexFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewPlayerIndexFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayerName_Encode(t *testing.T) {
	tests := []struct {
		name string
		n    PlayerName
		want string
	}{
		{
			name: "リンネン",
			n:    "リンネン",
			want: n0.encoded,
		},
		{
			name: "上海ニック",
			n:    "上海ニック",
			want: n1.encoded,
		},
		{
			name: "ekarute",
			n:    "ekarute",
			want: n2.encoded,
		},
		{
			name: "ノオチャベス",
			n:    "ノオチャベス",
			want: n3.encoded,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Encode(); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayers_Marshal(t *testing.T) {
	tests := []struct {
		name    string
		p       Players
		want    XmlElement
		wantErr bool
	}{
		{
			name: "四麻",
			p: Players{
				Player{
					Name: n0.original,
					Rate: 2054.10,
					Dan:  16,
					Sex:  "F",
				},
				Player{
					Name: n1.original,
					Rate: 2159.83,
					Dan:  16,
					Sex:  "M",
				},
				Player{
					Name: n2.original,
					Rate: 2159.83,
					Dan:  16,
					Sex:  "M",
				},
				Player{
					Name: n3.original,
					Rate: 2204.17,
					Dan:  17,
					Sex:  "M",
				},
			},
			want: XmlElement{
				Name: "UN",
				Attr: []XmlAttr{
					{
						Name:  "n0",
						Value: n0.encoded,
					},
					{
						Name:  "n1",
						Value: n1.encoded,
					},
					{
						Name:  "n2",
						Value: n2.encoded,
					},
					{
						Name:  "n3",
						Value: n3.encoded,
					},
					{
						Name:  "dan",
						Value: "16,16,16,17",
					},
					{
						Name:  "rate",
						Value: "2054.10,2159.83,2159.83,2204.17",
					},
					{
						Name:  "sx",
						Value: "F,M,M,M",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Marshal()
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 配列の順序まで一致させる必要があるので
			sort.Slice(got.Attr, func(i, j int) bool {
				return got.Attr[i].Name > got.Attr[j].Name
			})
			sort.Slice(tt.want.Attr, func(i, j int) bool {
				return tt.want.Attr[i].Name > tt.want.Attr[j].Name
			})
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Marshal() got = %v, want %v, diff %s", got, tt.want, diff)
			}
		})
	}
}

func TestPlayers_Unmarshal(t *testing.T) {
	type args struct {
		e XmlElement
	}
	tests := []struct {
		name    string
		p       Players
		args    args
		wantErr bool
	}{
		{
			name: "正常系:四麻",
			p: Players{
				{
					Name: "リンネン",
					Rate: 2054.10,
					Dan:  0,
					Sex:  "F",
				},
				{
					Name: "上海ニック",
					Rate: 2159.83,
					Dan:  0,
					Sex:  "M",
				},
				{
					Name: "ekarute",
					Rate: 2159.83,
					Dan:  0,
					Sex:  "M",
				},
				{
					Name: "ノオチャベス",
					Rate: 2204.17,
					Dan:  0,
					Sex:  "M",
				},
			},
			args: args{
				e: XmlElement{
					Name: "UN",
					Attr: []XmlAttr{
						{
							Name:  "n0",
							Value: "%E3%83%AA%E3%83%B3%E3%83%8D%E3%83%B3",
						},
						{
							Name:  "n1",
							Value: "%E4%B8%8A%E6%B5%B7%E3%83%8B%E3%83%83%E3%82%AF",
						},
						{
							Name:  "n2",
							Value: "%65%6B%61%72%75%74%65",
						},
						{
							Name:  "n3",
							Value: "%E3%83%8E%E3%82%AA%E3%83%81%E3%83%A3%E3%83%99%E3%82%B9",
						},
						{
							Name:  "dan",
							Value: "16,16,16,17",
						},
						{
							Name:  "rate",
							Value: "2054.10,2159.83,2159.83,2204.17",
						},
						{
							Name:  "sx",
							Value: "F,M,M,M",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "正常系:三麻",
			p: Players{
				{
					Name: "リンネン",
					Rate: 2054.10,
					Dan:  0,
					Sex:  "F",
				},
				{
					Name: "上海ニック",
					Rate: 2159.83,
					Dan:  0,
					Sex:  "M",
				},
				{
					Name: "ekarute",
					Rate: 2159.83,
					Dan:  0,
					Sex:  "M",
				},
			},
			args: args{
				e: XmlElement{
					Name: "UN",
					Attr: []XmlAttr{
						{
							Name:  "n0",
							Value: "%E3%83%AA%E3%83%B3%E3%83%8D%E3%83%B3",
						},
						{
							Name:  "n1",
							Value: "%E4%B8%8A%E6%B5%B7%E3%83%8B%E3%83%83%E3%82%AF",
						},
						{
							Name:  "n2",
							Value: "%65%6B%61%72%75%74%65",
						},
						{
							Name:  "n3",
							Value: "",
						},
						{
							Name:  "dan",
							Value: "16,16,16",
						},
						{
							Name:  "rate",
							Value: "2054.10,2159.83,2159.83",
						},
						{
							Name:  "sx",
							Value: "F,M,M",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "異常系:Danを整数値でSplitできない",
			p: Players{
				{
					Name: "リンネン",
					Rate: 2054.10,
					Dan:  0,
					Sex:  "F",
				},
				{
					Name: "上海ニック",
					Rate: 2159.83,
					Dan:  0,
					Sex:  "M",
				},
				{
					Name: "ekarute",
					Rate: 2159.83,
					Dan:  0,
					Sex:  "M",
				},
			},
			args: args{
				e: XmlElement{
					Name: "UN",
					Attr: []XmlAttr{
						{
							Name:  "n0",
							Value: "%E3%83%AA%E3%83%B3%E3%83%8D%E3%83%B3",
						},
						{
							Name:  "n1",
							Value: "%E4%B8%8A%E6%B5%B7%E3%83%8B%E3%83%83%E3%82%AF",
						},
						{
							Name:  "n2",
							Value: "%65%6B%61%72%75%74%65",
						},
						{
							Name:  "n3",
							Value: "",
						},
						{
							Name:  "dan",
							Value: "16,a,16",
						},
						{
							Name:  "rate",
							Value: "2054.10,2159.83,2159.83",
						},
						{
							Name:  "sx",
							Value: "F,M,M",
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Unmarshal(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
