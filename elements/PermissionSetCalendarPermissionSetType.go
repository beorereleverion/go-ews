package elements

// The PermissionSet element contains all the permissions that are configured for a calendar folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/permissionset-calendarpermissionsettype
import "encoding/xml"

type PermissionSetCalendarPermissionSetType struct {
	XMLName xml.Name

	// The CalendarPermissions element contains an array of calendar permissions for a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	CalendarPermissions *CalendarPermissions `xml:"CalendarPermissions"`
	// The UnknownEntries element contains an array of unknown permission entries that cannot be resolved against the Active Directory directory service. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	UnknownEntries *UnknownEntries `xml:"UnknownEntries"`
}

func (P *PermissionSetCalendarPermissionSetType) SetForMarshal() {
	P.XMLName.Local = "t:PermissionSet"
}

func (P *PermissionSetCalendarPermissionSetType) GetSchema() *Schema {
	return &SchemaTypes
}
