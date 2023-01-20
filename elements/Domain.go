package elements

// The Domain element identifies a single SMTP domain.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/domain
import "encoding/xml"

type Domain struct {
	XMLName xml.Name

	// Indicates whether subdomains of the domain identified by the Name attribute are considered internal. This attribute is optional.
	IncludeSubdomains *string `xml:"IncludeSubdomains,attr"`
	// Identifies the name of a domain. This attribute is required.
	Name *string `xml:"Name,attr"`
}

func (D *Domain) SetForMarshal() {
	D.XMLName.Local = "t:Domain"
}

func (D *Domain) GetSchema() *Schema {
	return &SchemaTypes
}
