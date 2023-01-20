package elements

// The Photo (MasterMailboxType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/photo-mastermailboxtype
import "encoding/xml"

type PhotoMasterMailboxType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
