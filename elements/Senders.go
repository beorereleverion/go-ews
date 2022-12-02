package elements

// The Senders element specifies an array of Simple Mail Transfer Protocol (SMTP) addresses.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/senders
type Senders struct {
	// The SmtpAddress element represents the Simple Mail Transfer Protocol (SMTP) address of an account to be used for impersonation or a Simple Mail Transfer Protocol (SMTP) recipient address of a calendar or contact sharing request.
	SmtpAddress *SmtpAddress `xml:"t:SmtpAddress"`
}
