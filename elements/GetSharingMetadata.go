package elements

// The GetSharingMetadata element defines a request to get an opaque authentication token that identifies the sharing invitation. This element is the base element for the GetSharingMetadata operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getsharingmetadata
import "encoding/xml"

type GetSharingMetadata struct {
	XMLName xml.Name

	// The IdOfFolderToShare element represents the identifier of the folder on the server that will be shared.
	IdOfFolderToShare *IdOfFolderToShare `xml:"IdOfFolderToShare"`
	// The Recipients element specifies an array of recipients of a message.
	Recipients *RecipientsArrayOfSmtpAddressType `xml:"Recipients"`
	// The SenderSmtpAddress element represents the SMTP e-mail address corresponding to the mailbox that contains the folder that will be shared.
	SenderSmtpAddress *SenderSmtpAddress `xml:"SenderSmtpAddress"`
}

func (G *GetSharingMetadata) SetForMarshal() {
	G.XMLName.Local = "m:GetSharingMetadata"
}

func (G *GetSharingMetadata) GetSchema() *Schema {
	return &SchemaMessages
}
