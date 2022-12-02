package elements

// The AddDistributionGroupToImList element defines a request to add a distribution list to an instant message list.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/adddistributiongrouptoimlist
type AddDistributionGroupToImList struct {
	// The DisplayName element contains the display name of a new instant messaging group contact or the display name of a new instant messaging group.
	DisplayName *DisplayNameNonEmptyStringType `xml:"t:DisplayName"`
	// The SmtpAddress element represents the Simple Mail Transfer Protocol (SMTP) address of an account to be used for impersonation or a Simple Mail Transfer Protocol (SMTP) recipient address of a calendar or contact sharing request.
	SmtpAddress *SmtpAddress `xml:"t:SmtpAddress"`
}
