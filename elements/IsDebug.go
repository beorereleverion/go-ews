package elements

// The IsDebug element is not used.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isdebug
import "encoding/xml"

type IsDebug struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsDebug) SetForMarshal() {
	I.XMLName.Local = "m:IsDebug"
}

func (I *IsDebug) GetSchema() *Schema {
	return &SchemaMessages
}
