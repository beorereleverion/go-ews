package elements

// The SharePointUrl (MasterMailboxType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sharepointurl-mastermailboxtype
import "encoding/xml"

type SharePointUrlMasterMailboxType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
