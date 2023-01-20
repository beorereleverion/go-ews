package elements

// The IsSigned element indicates whether incoming messages must be signed in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/issigned
import "encoding/xml"

type IsSigned struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsSignedfalse bool = false
	// true
	IsSignedtrue bool = true
)

func (I *IsSigned) SetForMarshal() {
	I.XMLName.Local = "m:IsSigned"
}

func (I *IsSigned) GetSchema() *Schema {
	return &SchemaMessages
}
