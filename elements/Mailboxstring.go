package elements

// The Mailbox element contains an identifier for a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailbox-string
import "encoding/xml"

type Mailboxstring struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}
