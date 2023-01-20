package elements

// The IsOwner element specifies whether the specified email user is the owner.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isowner
import "encoding/xml"

type IsOwner struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsOwnerfalse bool = false
	// true
	IsOwnertrue bool = true
)

func (I *IsOwner) SetForMarshal() {
	I.XMLName.Local = "t:IsOwner"
}

func (I *IsOwner) GetSchema() *Schema {
	return &SchemaTypes
}
