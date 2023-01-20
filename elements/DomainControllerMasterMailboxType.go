package elements

// The DomainController (MasterMailboxType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/domaincontroller-mastermailboxtype
import "encoding/xml"

type DomainControllerMasterMailboxType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
