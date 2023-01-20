package elements

// The ConnectingSID element represents an account to impersonate when you are using the ExchangeImpersonation SOAP header.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/connectingsid
import "encoding/xml"

type ConnectingSID struct {
	XMLName xml.Name

	// The PrimarySmtpAddress element represents the primary Simple Mail Transfer Protocol (SMTP) address of an account to be used for server-to-server authorization or delegate access.
	PrimarySmtpAddress *PrimarySmtpAddress `xml:"PrimarySmtpAddress"`
	// The PrincipalName element represents the user principal name (UPN) of the account to be used for Exchange impersonation.
	PrincipalName *PrincipalName `xml:"PrincipalName"`
	// The SID element represents the security descriptor definition language (SDDL) form of the security identifier (SID) for the account to use for impersonation or delegate access.
	SID *SID `xml:"SID"`
	// The SmtpAddress element represents the Simple Mail Transfer Protocol (SMTP) address of an account to be used for impersonation or a Simple Mail Transfer Protocol (SMTP) recipient address of a calendar or contact sharing request.
	SmtpAddress *SmtpAddress `xml:"SmtpAddress"`
}

func (C *ConnectingSID) SetForMarshal() {
	C.XMLName.Local = "t:ConnectingSID"
}

func (C *ConnectingSID) GetSchema() *Schema {
	return &SchemaTypes
}
