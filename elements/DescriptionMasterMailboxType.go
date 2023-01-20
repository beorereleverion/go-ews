package elements

// The Description (MasterMailboxType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/description-mastermailboxtype
import "encoding/xml"

type DescriptionMasterMailboxType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
