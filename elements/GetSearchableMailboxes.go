package elements

// The GetSearchableMailboxes element contains a request to get a list of mailboxes that the client has permission to perform an eDiscovery search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getsearchablemailboxes
import "encoding/xml"

type GetSearchableMailboxes struct {
	XMLName xml.Name

	// The ExpandGroupMembership element indicates whether to expand the membership of the group returned from a GetSearchableMailboxes request.
	ExpandGroupMembership *ExpandGroupMembership `xml:"ExpandGroupMembership"`
	// The SearchFilter element contains the query string to filter the mailboxes to be returned from a GetSearchableMailboxes request.
	SearchFilter *SearchFilter `xml:"SearchFilter"`
}

func (G *GetSearchableMailboxes) SetForMarshal() {
	G.XMLName.Local = "m:GetSearchableMailboxes"
}

func (G *GetSearchableMailboxes) GetSchema() *Schema {
	return &SchemaMessages
}
