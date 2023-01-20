package elements

// The GetUserOofSettingsRequest element is the root element that contains the arguments used to get a mailbox user's Out of Office (OOF) settings.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuseroofsettingsrequest
import "encoding/xml"

type GetUserOofSettingsRequest struct {
	XMLName xml.Name

	// The Mailbox element represents the mailbox user for a SetUserOofSettings or GetUserOofSettings request.
	Mailbox *MailboxAvailability `xml:"Mailbox"`
}

func (G *GetUserOofSettingsRequest) SetForMarshal() {
	G.XMLName.Local = "m:GetUserOofSettingsRequest"
}

func (G *GetUserOofSettingsRequest) GetSchema() *Schema {
	return &SchemaMessages
}
