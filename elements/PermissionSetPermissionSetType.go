package elements

// The PermissionSet element contains all the permissions that are configured for a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/permissionset-permissionsettype
import "encoding/xml"

type PermissionSetPermissionSetType struct {
	XMLName xml.Name

	// The Permissions element contains the collection of permissions for a folder.
	Permissions *Permissions `xml:"Permissions"`
	// The UnknownEntries element contains an array of unknown permission entries that cannot be resolved against the Active Directory directory service. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	UnknownEntries *UnknownEntries `xml:"UnknownEntries"`
}

func (P *PermissionSetPermissionSetType) SetForMarshal() {
	P.XMLName.Local = "t:PermissionSet"
}

func (P *PermissionSetPermissionSetType) GetSchema() *Schema {
	return &SchemaTypes
}
