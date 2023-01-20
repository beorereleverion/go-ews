package elements

// The InternalDomains element identifies the list of internal SMTP domains of the organization.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/internaldomains-smtpdomainlist
import "encoding/xml"

type InternalDomainsSmtpDomainList struct {
	XMLName xml.Name

	// The Domain element identifies a single SMTP domain.
	Domain *Domain `xml:"Domain"`
}

func (I *InternalDomainsSmtpDomainList) SetForMarshal() {
	I.XMLName.Local = "t:InternalDomains"
}

func (I *InternalDomainsSmtpDomainList) GetSchema() *Schema {
	return &SchemaTypes
}
