package elements

// The FieldValue element represents the value of the field that caused the validation error.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fieldvalue
import "encoding/xml"

type FieldValue struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (F *FieldValue) SetForMarshal() {
	F.XMLName.Local = "m:FieldValue"
}

func (F *FieldValue) GetSchema() *Schema {
	return &SchemaMessages
}
