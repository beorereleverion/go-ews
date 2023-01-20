package elements

// The ApplicationRoles element specifies the application roles that the calling partner application uses for the current call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/applicationroles
import "encoding/xml"

type ApplicationRoles struct {
	XMLName xml.Name

	// The Role element specifies a string that represents a management role.
	Role *Role `xml:"Role"`
}

func (A *ApplicationRoles) SetForMarshal() {
	A.XMLName.Local = "t:ApplicationRoles"
}

func (A *ApplicationRoles) GetSchema() *Schema {
	return &SchemaTypes
}
