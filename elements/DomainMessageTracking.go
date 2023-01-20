package elements

// The Domain element represents the domain to search for.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/domain-message-tracking
import "encoding/xml"

type DomainMessageTracking struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DomainMessageTracking) SetForMarshal() {
	D.XMLName.Local = "m:Domain"
}

func (D *DomainMessageTracking) GetSchema() *Schema {
	return &SchemaMessages
}
