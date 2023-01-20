package elements

// The GetFederatedDirectoryGroup element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getfederateddirectorygroup
import "encoding/xml"

type GetFederatedDirectoryGroup struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
