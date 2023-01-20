package elements

// The GetFederatedDirectoryGroupResponse element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getfederateddirectorygroupresponse
import "encoding/xml"

type GetFederatedDirectoryGroupResponse struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
