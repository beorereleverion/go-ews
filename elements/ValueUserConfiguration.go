package elements

// The Value element specifies the dictionary object value as a string.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/value-userconfiguration
import "encoding/xml"

type ValueUserConfiguration struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (V *ValueUserConfiguration) SetForMarshal() {
	V.XMLName.Local = "t:Value"
}

func (V *ValueUserConfiguration) GetSchema() *Schema {
	return &SchemaTypes
}
