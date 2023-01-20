package elements

// The UpdateMailboxAssociation element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updatemailboxassociation
import "encoding/xml"

type UpdateMailboxAssociation struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
