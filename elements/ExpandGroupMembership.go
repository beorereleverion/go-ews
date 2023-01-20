package elements

// The ExpandGroupMembership element indicates whether to expand the membership of the group returned from a GetSearchableMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/expandgroupmembership
import "encoding/xml"

type ExpandGroupMembership struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	ExpandGroupMembershipfalse bool = false
	// true
	ExpandGroupMembershiptrue bool = true
)

func (E *ExpandGroupMembership) SetForMarshal() {
	E.XMLName.Local = "m:ExpandGroupMembership"
}

func (E *ExpandGroupMembership) GetSchema() *Schema {
	return &SchemaMessages
}
