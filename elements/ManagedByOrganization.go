package elements

// The ManagedByOrganization element specifies the managing organization.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/managedbyorganization
import "encoding/xml"

type ManagedByOrganization struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *ManagedByOrganization) SetForMarshal() {
	M.XMLName.Local = "t:ManagedByOrganization"
}

func (M *ManagedByOrganization) GetSchema() *Schema {
	return &SchemaTypes
}
