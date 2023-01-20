package elements

// The Permissions element contains the collection of permissions for a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/permissions
import "encoding/xml"

type Permissions struct {
	XMLName xml.Name

	// The Permission element defines the access that a user has to a folder.
	Permission *Permission `xml:"Permission"`
}

func (P *Permissions) SetForMarshal() {
	P.XMLName.Local = "t:Permissions"
}

func (P *Permissions) GetSchema() *Schema {
	return &SchemaTypes
}
