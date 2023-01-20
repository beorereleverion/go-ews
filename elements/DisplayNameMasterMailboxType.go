package elements

// The DisplayName (MasterMailboxType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displayname-mastermailboxtype
import "encoding/xml"

type DisplayNameMasterMailboxType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
