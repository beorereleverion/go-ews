package elements

// The ExternalDirectoryObjectId (FederatedDirectoryGroupType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/externaldirectoryobjectid-federateddirectorygrouptype
import "encoding/xml"

type ExternalDirectoryObjectIdFederatedDirectoryGroupType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
