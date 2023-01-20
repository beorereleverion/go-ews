package elements

// The GetFederatedDirectoryUser element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getfederateddirectoryuser
import "encoding/xml"

type GetFederatedDirectoryUser struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
