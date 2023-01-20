package elements

// The ConvertHtmlCodePageToUTF8 element indicates whether the item HTML body is converted to UTF8.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/converthtmlcodepagetoutf8
import "encoding/xml"

type ConvertHtmlCodePageToUTF8 struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (C *ConvertHtmlCodePageToUTF8) SetForMarshal() {
	C.XMLName.Local = "t:ConvertHtmlCodePageToUTF8"
}

func (C *ConvertHtmlCodePageToUTF8) GetSchema() *Schema {
	return &SchemaTypes
}
