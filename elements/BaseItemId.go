package elements

// The BaseItemId element represents the base class for IDs that represent items in a mailbox. This is an abstract class and therefore will not occur in an instance document.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/baseitemid
import "encoding/xml"

type BaseItemId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (B *BaseItemId) SetForMarshal() {
	B.XMLName.Local = "t:BaseItemId"
}

func (B *BaseItemId) GetSchema() *Schema {
	return &SchemaTypes
}
