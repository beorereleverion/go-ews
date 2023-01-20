package elements

// The GetDiscoverySearchConfiguration element specifies a request to retrieve the eDiscovery search configuration.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getdiscoverysearchconfiguration
import "encoding/xml"

type GetDiscoverySearchConfiguration struct {
	XMLName xml.Name

	// The ExpandGroupMembership element indicates whether to expand the membership of the group returned from a GetSearchableMailboxes request.
	ExpandGroupMembership *ExpandGroupMembership `xml:"ExpandGroupMembership"`
	// The SearchId element specifies the identifier of a discovery search.
	SearchId *SearchId `xml:"SearchId"`
}

func (G *GetDiscoverySearchConfiguration) SetForMarshal() {
	G.XMLName.Local = "m:GetDiscoverySearchConfiguration"
}

func (G *GetDiscoverySearchConfiguration) GetSchema() *Schema {
	return &SchemaMessages
}
