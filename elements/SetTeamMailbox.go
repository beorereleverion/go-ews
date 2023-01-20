package elements

// The SetTeamMailbox element contains a request to set a site mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setteammailbox
import "encoding/xml"

type SetTeamMailbox struct {
	XMLName xml.Name

	// The EmailAddress element specifies the fully resolved SMTP address for the site mailbox or the associated persona.
	EmailAddress *EmailAddressEmailAddressType `xml:"EmailAddress"`
	// The SharePointSiteURL element contains the Uniform Resource Locator (URL) of the SharePoint site that is linked with the site mailbox.
	SharePointSiteUrl *SharePointSiteUrl `xml:"SharePointSiteUrl"`
	// The State element contains the lifecycle state that is set on a site mailbox.
	State *StateTeamMailboxLifecycleStateType `xml:"State"`
}

func (S *SetTeamMailbox) SetForMarshal() {
	S.XMLName.Local = "m:SetTeamMailbox"
}

func (S *SetTeamMailbox) GetSchema() *Schema {
	return &SchemaMessages
}
