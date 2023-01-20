package elements

// The InboxUrl (FederatedDirectoryGroupType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/inboxurl-federateddirectorygrouptype
import "encoding/xml"

type InboxUrlFederatedDirectoryGroupType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
