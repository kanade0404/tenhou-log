package xml

import (
	"github.com/google/go-cmp/cmp"
	"sort"
	"testing"
)

func TestGameInfo_Marshal(t *testing.T) {
	type fields struct {
		Demo          bool
		Red           bool
		KuitanAtozuke bool
		East          bool
		Three         bool
		HighSpeed     bool
		Table         Table
		Lobby         string
	}
	tests := []struct {
		name    string
		fields  fields
		want    XmlElement
		wantErr bool
	}{
		{
			name: "四鳳南喰赤",
			fields: fields{
				Demo:          false,
				Red:           true,
				KuitanAtozuke: true,
				East:          false,
				Three:         false,
				HighSpeed:     false,
				Table:         Houou,
				Lobby:         "0",
			},
			want: XmlElement{
				Name: "GO",
				Attr: []XmlAttr{
					{
						Name:  "type",
						Value: "169",
					},
					{
						Name:  "lobby",
						Value: "0",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "四特南喰赤",
			fields: fields{
				Demo:          false,
				Red:           true,
				KuitanAtozuke: true,
				East:          false,
				Three:         false,
				HighSpeed:     false,
				Table:         Tokujou,
				Lobby:         "0",
			},
			want: XmlElement{
				Name: "GO",
				Attr: []XmlAttr{
					{
						Name:  "type",
						Value: "41",
					},
					{
						Name:  "lobby",
						Value: "0",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "三上東速",
			fields: fields{
				Demo:          false,
				Red:           false,
				KuitanAtozuke: false,
				East:          true,
				Three:         true,
				HighSpeed:     true,
				Table:         Jyoukyuu,
				Lobby:         "0",
			},
			want: XmlElement{
				Name: "GO",
				Attr: []XmlAttr{
					{
						Name:  "type",
						Value: "215",
					},
					{
						Name:  "lobby",
						Value: "0",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GameInfo{
				Demo:          tt.fields.Demo,
				Red:           tt.fields.Red,
				KuitanAtozuke: tt.fields.KuitanAtozuke,
				East:          tt.fields.East,
				Three:         tt.fields.Three,
				HighSpeed:     tt.fields.HighSpeed,
				Table:         tt.fields.Table,
				Lobby:         tt.fields.Lobby,
			}
			got, err := g.Marshal()
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			sort.Slice(tt.want.Attr, func(i, j int) bool {
				return tt.want.Attr[i].Name > tt.want.Attr[j].Name
			})
			sort.Slice(got.Attr, func(i, j int) bool {
				return got.Attr[i].Name > got.Attr[j].Name
			})
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Marshal() got = %v, want %v, diff: %s", got, tt.want, diff)
			}
		})
	}
}

func TestGameInfo_Unmarshal(t *testing.T) {
	type fields struct {
		Demo          bool
		Red           bool
		KuitanAtozuke bool
		East          bool
		Three         bool
		HighSpeed     bool
		Table         Table
		Lobby         string
	}
	type args struct {
		x XmlElement
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"四鳳南喰赤",
			fields{
				Demo:          false,
				Red:           true,
				KuitanAtozuke: true,
				East:          false,
				Three:         false,
				HighSpeed:     false,
				Table:         Houou,
				Lobby:         "0",
			},
			args{
				newXmlElement("GO", NewXmlAttrs(map[string]string{"type": "169", "lobby": "0"})),
			},
			false,
		},
		{
			"四特南喰赤",
			fields{
				Demo:          false,
				Red:           true,
				KuitanAtozuke: true,
				East:          false,
				Three:         false,
				HighSpeed:     false,
				Table:         Tokujou,
				Lobby:         "0",
			},
			args{
				newXmlElement("GO", NewXmlAttrs(map[string]string{"type": "41", "lobby": "0"})),
			},
			false,
		},
		{
			"四般南喰赤",
			fields{
				Demo:          false,
				Red:           true,
				KuitanAtozuke: false,
				East:          false,
				Three:         false,
				HighSpeed:     false,
				Table:         Ippan,
				Lobby:         "0",
			},
			args{
				newXmlElement("GO", NewXmlAttrs(map[string]string{"type": "9", "lobby": "0"})),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GameInfo{
				Demo:          tt.fields.Demo,
				Red:           tt.fields.Red,
				KuitanAtozuke: tt.fields.KuitanAtozuke,
				East:          tt.fields.East,
				Three:         tt.fields.Three,
				HighSpeed:     tt.fields.HighSpeed,
				Table:         tt.fields.Table,
				Lobby:         tt.fields.Lobby,
			}
			if err := g.Unmarshal(tt.args.x); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func TestTable_Name(t *testing.T) {
	tests := []struct {
		name string
		t    Table
		want string
	}{
		{
			name: "一般卓",
			t:    Table(0x00),
			want: "一般",
		},
		{
			name: "上級卓",
			t:    Table(0x80),
			want: "上級",
		},
		{
			name: "特上卓",
			t:    Table(0x20),
			want: "特上",
		},
		{
			name: "鳳凰卓",
			t:    Table(0xA0),
			want: "鳳凰",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Name(); got != tt.want {
				t.Errorf("File() = %v, want %v", got, tt.want)
			}
		})
	}
}
