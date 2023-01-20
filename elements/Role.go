package elements

// The Role element specifies a string that represents a management role.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/role
import "encoding/xml"

type Role struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (R *Role) SetForMarshal() {
	R.XMLName.Local = "t:Role"
}

func (R *Role) GetSchema() *Schema {
	return &SchemaTypes
}
