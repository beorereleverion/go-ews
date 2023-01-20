package elements

// The UnpinTeamMailbox element contains the request to unpin a site mailbox from the client by removing it from the Autodiscover response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/unpinteammailbox
import "encoding/xml"

type UnpinTeamMailbox struct {
	XMLName xml.Name

	// The EmailAddress element specifies the fully resolved SMTP address for the site mailbox or the associated persona.
	EmailAddress *EmailAddressEmailAddressType `xml:"EmailAddress"`
}

func (U *UnpinTeamMailbox) SetForMarshal() {
	U.XMLName.Local = "m:UnpinTeamMailbox"
}

func (U *UnpinTeamMailbox) GetSchema() *Schema {
	return &SchemaMessages
}
