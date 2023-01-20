package elements

// The ManagementRole element specifies a list of user and application management roles.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/managementrole
import "encoding/xml"

type ManagementRole struct {
	XMLName xml.Name

	// The ApplicationRoles element specifies the application roles that the calling partner application uses for the current call.
	ApplicationRoles *ApplicationRoles `xml:"ApplicationRoles"`
	// The UserRoles element specifies the user roles that the calling user, or the user that the calling partner application is acting as, wants to apply to the current call.
	UserRoles *UserRoles `xml:"UserRoles"`
}

func (M *ManagementRole) SetForMarshal() {
	M.XMLName.Local = "t:ManagementRole"
}

func (M *ManagementRole) GetSchema() *Schema {
	return &SchemaTypes
}
