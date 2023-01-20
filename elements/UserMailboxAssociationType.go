package elements

// The User (MailboxAssociationType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/user-mailboxassociationtype
import "encoding/xml"

type UserMailboxAssociationType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
