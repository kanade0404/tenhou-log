package xml

import "fmt"

type XmlElement struct {
	Name string
	Attr []XmlAttr
}

func (x XmlElement) Text() string {
	var text string
	for _, attr := range x.Attr {
		text += fmt.Sprintf(` %s="%s"`, attr.Name, attr.Value)
	}
	return text
}

func (x *XmlElement) AppendAttr(name, value string) {
	x.Attr = append(x.Attr, XmlAttr{
		Name:  name,
		Value: value,
	})
}

func (x XmlElement) AttrByName(name string) string {
	for _, attr := range x.Attr {
		if attr.Name == name {
			return attr.Value
		}
	}
	return ""
}

func newXmlElement(name string, attr []XmlAttr) XmlElement {
	el := XmlElement{
		Name: name,
	}
	for _, at := range attr {
		el.Attr = append(el.Attr, XmlAttr{
			Name:  at.Name,
			Value: at.Value,
		})
	}
	return el
}

func (x *XmlElement) ForEachAttr(f func(name, value string) error) error {
	for _, attr := range x.Attr {
		if err := f(attr.Name, attr.Value); err != nil {
			return err
		}
	}
	return nil
}

type XmlAttr struct {
	Name  string
	Value string
}

func NewXmlAttr(name, value string) XmlAttr {
	return XmlAttr{
		Name:  name,
		Value: value,
	}
}

func NewXmlAttrs(attrs map[string]string) []XmlAttr {
	var result []XmlAttr
	for k, v := range attrs {
		result = append(result, NewXmlAttr(k, v))
	}
	return result
}

// Mjloggm XML全体のタグ
type Mjloggm struct {
}

// SHUFFLE 乱数情報
type SHUFFLE struct {
	Seed string `xml:"seed,attr"`
}

type UN struct {
	Type byte `xml:"type,attr"`
}
