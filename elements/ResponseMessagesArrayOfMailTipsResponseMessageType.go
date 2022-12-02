package elements

// The ResponseMessages element represents a list of mail tips response messages.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responsemessages-arrayofmailtipsresponsemessagetype
type ResponseMessagesArrayOfMailTipsResponseMessageType struct {
	// The MailTipsResponseMessageType element represents mail tips settings.
	MailTipsResponseMessageType *MailTipsResponseMessageType `xml:"m:MailTipsResponseMessageType"`
}
