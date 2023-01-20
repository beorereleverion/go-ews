package elements

// The GetFederatedDirectoryUserResponse element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getfederateddirectoryuserresponse
import "encoding/xml"

type GetFederatedDirectoryUserResponse struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
