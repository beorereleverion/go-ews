package elements

// The InboxUrl (MasterMailboxType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/inboxurl-mastermailboxtype
import "encoding/xml"

type InboxUrlMasterMailboxType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
