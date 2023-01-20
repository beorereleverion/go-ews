package elements

// The GetMailTips element represents the recipients and types of mail tips to retrieve.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getmailtips
import "encoding/xml"

type GetMailTips struct {
	XMLName xml.Name

	// The MailTipsRequested element contains the types of mail tips requested from the service.
	MailTipsRequested *MailTipsRequested `xml:"MailTipsRequested"`
	// The Recipients element represents a collection of recipients that receive a copy of the message.
	Recipients *RecipientsArrayOfRecipientsType `xml:"Recipients"`
	// The SendingAs element represents an e-mail address that a user is trying to send as.
	SendingAs *SendingAs `xml:"SendingAs"`
}

func (G *GetMailTips) SetForMarshal() {
	G.XMLName.Local = "m:GetMailTips"
}

func (G *GetMailTips) GetSchema() *Schema {
	return &SchemaMessages
}
