package elements

// The IsDelegated element indicates whether a meeting was handled by an account that has delegate access.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isdelegated
import "encoding/xml"

type IsDelegated struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsDelegated) SetForMarshal() {
	I.XMLName.Local = "t:IsDelegated"
}

func (I *IsDelegated) GetSchema() *Schema {
	return &SchemaTypes
}
