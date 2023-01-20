package elements

// The IsQuickContact element specifies a Boolean value that indicates whether the underlying contact is a quick contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isquickcontact
import "encoding/xml"

type IsQuickContact struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsQuickContactfalse bool = false
	// true
	IsQuickContacttrue bool = true
)

func (I *IsQuickContact) SetForMarshal() {
	I.XMLName.Local = "t:IsQuickContact"
}

func (I *IsQuickContact) GetSchema() *Schema {
	return &SchemaTypes
}
