package elements

// The UpdateMailboxAssociationResponse element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updatemailboxassociationresponse
import "encoding/xml"

type UpdateMailboxAssociationResponse struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
