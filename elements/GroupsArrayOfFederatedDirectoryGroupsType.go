package elements

// The Groups (ArrayOfFederatedDirectoryGroupsType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/groups-arrayoffederateddirectorygroupstype
import "encoding/xml"

type GroupsArrayOfFederatedDirectoryGroupsType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
