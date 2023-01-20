package elements

// The ExternalDirectoryObjectId element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/externaldirectoryobjectid
import "encoding/xml"

type ExternalDirectoryObjectId struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
