package elements

// The DiscoverySearchConfiguration element specifies the configuration for eDiscovery search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/discoverysearchconfiguration
import "encoding/xml"

type DiscoverySearchConfiguration struct {
	XMLName xml.Name

	// The SearchId element specifies the identifier of a discovery search.
	SearchId *SearchId `xml:"SearchId"`
	// The SearchQuery element specifies the discovery search query.
	SearchQuery *SearchQuery `xml:"SearchQuery"`
	// The SearchableMailboxes element contains an array of the mailboxes returned from a GetSearchableMailboxes request.
	SearchableMailboxes *SearchableMailboxes `xml:"SearchableMailboxes"`
}

func (D *DiscoverySearchConfiguration) SetForMarshal() {
	D.XMLName.Local = "m:DiscoverySearchConfiguration"
}

func (D *DiscoverySearchConfiguration) GetSchema() *Schema {
	return &SchemaMessages
}
