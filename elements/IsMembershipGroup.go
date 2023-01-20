package elements

// The IsMembershipGroup element specifies a Boolean value that indicates whether the entity is a distribution group or a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ismembershipgroup
import "encoding/xml"

type IsMembershipGroup struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsMembershipGroupfalse bool = false
	// true
	IsMembershipGrouptrue bool = true
)

func (I *IsMembershipGroup) SetForMarshal() {
	I.XMLName.Local = "t:IsMembershipGroup"
}

func (I *IsMembershipGroup) GetSchema() *Schema {
	return &SchemaTypes
}
