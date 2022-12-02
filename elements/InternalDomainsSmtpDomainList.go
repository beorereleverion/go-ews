package elements

// The InternalDomains element identifies the list of internal SMTP domains of the organization.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/internaldomains-smtpdomainlist
type InternalDomainsSmtpDomainList struct {
	// The Domain element identifies a single SMTP domain.
	Domain *Domain `xml:"t:Domain"`
}
