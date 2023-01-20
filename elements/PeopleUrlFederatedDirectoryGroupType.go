package elements

// The PeopleUrl (FederatedDirectoryGroupType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/peopleurl-federateddirectorygrouptype
import "encoding/xml"

type PeopleUrlFederatedDirectoryGroupType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
