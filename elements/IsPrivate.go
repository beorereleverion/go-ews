package elements

// The IsPrivate element indicates whether the calendar item is private.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isprivate
import "encoding/xml"

type IsPrivate struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsPrivatefalse bool = false
	// true
	IsPrivatetrue bool = true
)

func (I *IsPrivate) SetForMarshal() {
	I.XMLName.Local = "t:IsPrivate"
}

func (I *IsPrivate) GetSchema() *Schema {
	return &SchemaTypes
}
