package elements

// The PhotoUrl (FederatedDirectoryGroupType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/photourl-federateddirectorygrouptype
import "encoding/xml"

type PhotoUrlFederatedDirectoryGroupType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
