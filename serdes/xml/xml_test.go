package xml

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestNewXmlAttr(t *testing.T) {
	type args struct {
		name  string
		value string
	}
	tests := []struct {
		name string
		args args
		want XmlAttr
	}{
		{
			name: "正常系",
			args: args{
				name:  "hoge",
				value: "huga",
			},
			want: XmlAttr{
				Name:  "hoge",
				Value: "huga",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewXmlAttr(tt.args.name, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewXmlAttr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXmlElement_AppendAttr(t *testing.T) {
	type fields struct {
		Name string
		Attr []XmlAttr
	}
	type args struct {
		name  string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   XmlElement
	}{
		{
			name: "正常系",
			fields: fields{
				Name: "test",
				Attr: []XmlAttr{
					{
						Name:  "hoge",
						Value: "huga",
					},
				},
			},
			args: args{
				name:  "foo",
				value: "bar",
			},
			want: XmlElement{
				Name: "test",
				Attr: []XmlAttr{
					{
						Name:  "hoge",
						Value: "huga",
					},
					{
						Name:  "foo",
						Value: "bar",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := XmlElement{
				Name: tt.fields.Name,
				Attr: tt.fields.Attr,
			}
			x.AppendAttr(tt.args.name, tt.args.value)
			if diff := cmp.Diff(x, tt.want); diff != "" {
				t.Errorf("diff: %s, actual:%v, want:%v", diff, x, tt.want)
			}
		})
	}
}

func TestXmlElement_AttrByName(t *testing.T) {
	type fields struct {
		Name string
		Attr []XmlAttr
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "正常系",
			fields: fields{
				Name: "test",
				Attr: []XmlAttr{
					{
						Name:  "hoge",
						Value: "huga",
					},
					{
						Name:  "foo",
						Value: "bar",
					},
				},
			},
			args: args{
				name: "hoge",
			},
			want: "huga",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := XmlElement{
				Name: tt.fields.Name,
				Attr: tt.fields.Attr,
			}
			if got := x.AttrByName(tt.args.name); got != tt.want {
				t.Errorf("AttrByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXmlElement_Text(t *testing.T) {
	type fields struct {
		Name string
		Attr []XmlAttr
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "正常系",
			fields: fields{
				Name: "test",
				Attr: []XmlAttr{
					{
						Name:  "hoge",
						Value: "huga",
					},
					{
						Name:  "foo",
						Value: "bar",
					},
				},
			},
			want: ` hoge="huga" foo="bar"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := XmlElement{
				Name: tt.fields.Name,
				Attr: tt.fields.Attr,
			}
			if got := x.Text(); got != tt.want {
				t.Errorf("Text() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newXmlElement(t *testing.T) {
	type args struct {
		name string
		attr []XmlAttr
	}
	tests := []struct {
		name string
		args args
		want XmlElement
	}{
		{
			name: "正常系",
			args: args{
				name: "test",
				attr: []XmlAttr{
					{
						Name:  "hoge",
						Value: "huga",
					},
					{
						Name:  "foo",
						Value: "bar",
					},
				},
			},
			want: XmlElement{
				Name: "test",
				Attr: []XmlAttr{
					{
						Name:  "hoge",
						Value: "huga",
					},
					{
						Name:  "foo",
						Value: "bar",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newXmlElement(tt.args.name, tt.args.attr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newXmlElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
