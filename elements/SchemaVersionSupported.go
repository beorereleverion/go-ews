package elements

// The SchemaVersionSupported element contains the version of the manifest schema supported by the client.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/schemaversionsupported
import "encoding/xml"

type SchemaVersionSupported struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SchemaVersionSupported) SetForMarshal() {
	S.XMLName.Local = "m:SchemaVersionSupported"
}

func (S *SchemaVersionSupported) GetSchema() *Schema {
	return &SchemaMessages
}
