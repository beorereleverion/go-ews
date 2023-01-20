package elements

// The Group (MailboxAssociationType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/group-mailboxassociationtype
import "encoding/xml"

type GroupMailboxAssociationType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
