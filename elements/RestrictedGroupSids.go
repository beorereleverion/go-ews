package elements

// The RestrictedGroupSids element represents a collection of restricted groups from a user's token.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/restrictedgroupsids
import "encoding/xml"

type RestrictedGroupSids struct {
	XMLName xml.Name

	// The RestrictGroupIdentifier element represents the group security identifier (SID) and attributes for a restricted group within a user token.
	RestrictedGroupIdentifier *RestrictedGroupIdentifier `xml:"RestrictedGroupIdentifier"`
}

func (R *RestrictedGroupSids) SetForMarshal() {
	R.XMLName.Local = "t:RestrictedGroupSids"
}

func (R *RestrictedGroupSids) GetSchema() *Schema {
	return &SchemaTypes
}
