package elements

// The SharingEffectiveRights element indicates the permissions that the user has for the contact data that is being shared.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sharingeffectiverights-permissionreadaccesstype
import "encoding/xml"

type SharingEffectiveRightsPermissionReadAccessType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// FullDetails
	SharingEffectiveRightsPermissionReadAccessTypeFullDetails string = `FullDetails`
	// None
	SharingEffectiveRightsPermissionReadAccessTypeNone string = `None`
)

func (S *SharingEffectiveRightsPermissionReadAccessType) SetForMarshal() {
	S.XMLName.Local = "t:SharingEffectiveRights"
}

func (S *SharingEffectiveRightsPermissionReadAccessType) GetSchema() *Schema {
	return &SchemaTypes
}
