package elements

// The SmtpAddress (MasterMailboxType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/smtpaddress-mastermailboxtype
import "encoding/xml"

type SmtpAddressMasterMailboxType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
