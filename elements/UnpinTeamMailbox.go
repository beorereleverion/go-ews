package elements

// The UnpinTeamMailbox element contains the request to unpin a site mailbox from the client by removing it from the Autodiscover response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/unpinteammailbox
type UnpinTeamMailbox struct {
	// The EmailAddress element specifies the fully resolved SMTP address for the site mailbox or the associated persona.
	EmailAddress *EmailAddressEmailAddressType `xml:"t:EmailAddress"`
}
