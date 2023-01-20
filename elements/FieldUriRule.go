package elements

// The FieldURI element specifies the URI to the rule field that caused the validation error.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fielduri-rule
import "encoding/xml"

type FieldUriRule struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (F *FieldUriRule) SetForMarshal() {
	F.XMLName.Local = "m:FieldUri"
}

func (F *FieldUriRule) GetSchema() *Schema {
	return &SchemaMessages
}
