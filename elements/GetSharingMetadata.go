package elements

// The GetSharingMetadata element defines a request to get an opaque authentication token that identifies the sharing invitation. This element is the base element for the GetSharingMetadata operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getsharingmetadata
type GetSharingMetadata struct {
	// The IdOfFolderToShare element represents the identifier of the folder on the server that will be shared.
	IdOfFolderToShare *IdOfFolderToShare `xml:"m:IdOfFolderToShare"`
	// The Recipients element specifies an array of recipients of a message.
	Recipients *RecipientsArrayOfSmtpAddressType `xml:"m:Recipients"`
	// The SenderSmtpAddress element represents the SMTP e-mail address corresponding to the mailbox that contains the folder that will be shared.
	SenderSmtpAddress *SenderSmtpAddress `xml:"m:SenderSmtpAddress"`
}
