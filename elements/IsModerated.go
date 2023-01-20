package elements

// The IsModerated element indicates whether the recipient's mailbox is being moderated.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ismoderated
import "encoding/xml"

type IsModerated struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsModeratedfalse bool = false
	// true
	IsModeratedtrue bool = true
)

func (I *IsModerated) SetForMarshal() {
	I.XMLName.Local = "t:IsModerated"
}

func (I *IsModerated) GetSchema() *Schema {
	return &SchemaTypes
}
