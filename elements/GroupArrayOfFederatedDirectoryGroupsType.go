package elements

// The Group (ArrayOfFederatedDirectoryGroupsType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/group-arrayoffederateddirectorygroupstype
import "encoding/xml"

type GroupArrayOfFederatedDirectoryGroupsType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
