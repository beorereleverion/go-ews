package elements

// The SetUserOofSettingsRequest element contains the arguments used to set a mailbox user's Out of Office (OOF) settings.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setuseroofsettingsrequest
import "encoding/xml"

type SetUserOofSettingsRequest struct {
	XMLName xml.Name

	// The Mailbox element represents the mailbox user for a SetUserOofSettings or GetUserOofSettings request.
	Mailbox *MailboxAvailability `xml:"Mailbox"`
	// The UserOofSettings element specifies the Out of Office (OOF) settings.
	UserOofSettings *UserOofSettings `xml:"UserOofSettings"`
}

func (S *SetUserOofSettingsRequest) SetForMarshal() {
	S.XMLName.Local = "m:SetUserOofSettingsRequest"
}

func (S *SetUserOofSettingsRequest) GetSchema() *Schema {
	return &SchemaMessages
}
