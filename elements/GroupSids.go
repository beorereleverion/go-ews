package elements

// The GroupSids element represents a collection of Active Directory directory service group object security identifiers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/groupsids
import "encoding/xml"

type GroupSids struct {
	XMLName xml.Name

	// The GroupIdentifier element represents a single security identifier and attribute for an Active Directory directory service object group of which the account is a member.
	GroupIdentifier *GroupIdentifier `xml:"GroupIdentifier"`
}

func (G *GroupSids) SetForMarshal() {
	G.XMLName.Local = "t:GroupSids"
}

func (G *GroupSids) GetSchema() *Schema {
	return &SchemaTypes
}
