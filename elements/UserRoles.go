package elements

// The UserRoles element specifies the user roles that the calling user, or the user that the calling partner application is acting as, wants to apply to the current call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/userroles
import "encoding/xml"

type UserRoles struct {
	XMLName xml.Name

	// The Role element specifies a string that represents a management role.
	Role *Role `xml:"Role"`
}

func (U *UserRoles) SetForMarshal() {
	U.XMLName.Local = "t:UserRoles"
}

func (U *UserRoles) GetSchema() *Schema {
	return &SchemaTypes
}
