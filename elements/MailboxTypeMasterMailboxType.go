package elements

// The MailboxType (MasterMailboxType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxtype-mastermailboxtype
import "encoding/xml"

type MailboxTypeMasterMailboxType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
