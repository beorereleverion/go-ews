package elements

// The ProtectionRulesConfiguration element contains service configuration information for the protection rules service.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/protectionrulesconfiguration
import "encoding/xml"

type ProtectionRulesConfiguration struct {
	XMLName xml.Name

	// The InternalDomains element identifies the list of internal SMTP domains of the organization.
	InternalDomains *InternalDomainsSmtpDomainList `xml:"InternalDomains"`
	// The Rules element contains an array of protection rules.
	Rules *Rules `xml:"Rules"`
	// Specifies how often, in whole hours, the client should request protection rules from the server. This attribute is required and its value must be an integer that is equal to or greater than 1.
	RefreshInterval *string `xml:"RefreshInterval,attr"`
}

func (P *ProtectionRulesConfiguration) SetForMarshal() {
	P.XMLName.Local = "m:ProtectionRulesConfiguration"
}

func (P *ProtectionRulesConfiguration) GetSchema() *Schema {
	return &SchemaMessages
}
