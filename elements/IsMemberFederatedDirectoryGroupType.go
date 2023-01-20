package elements

// The IsMember (FederatedDirectoryGroupType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ismember-federateddirectorygrouptype
import "encoding/xml"

type IsMemberFederatedDirectoryGroupType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
