package elements

// The Recipients element specifies an array of recipients of a message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recipients-arrayofsmtpaddresstype
type RecipientsArrayOfSmtpAddressType struct {
	// The SmtpAddress element represents the Simple Mail Transfer Protocol (SMTP) address of an account to be used for impersonation or a Simple Mail Transfer Protocol (SMTP) recipient address of a calendar or contact sharing request.
	SmtpAddress *SmtpAddress `xml:"t:SmtpAddress"`
}
