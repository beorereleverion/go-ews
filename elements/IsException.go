package elements

// The IsException element indicates whether an instance of a recurring calendar item is changed from the master.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isexception
import "encoding/xml"

type IsException struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsException) SetForMarshal() {
	I.XMLName.Local = "t:IsException"
}

func (I *IsException) GetSchema() *Schema {
	return &SchemaTypes
}
