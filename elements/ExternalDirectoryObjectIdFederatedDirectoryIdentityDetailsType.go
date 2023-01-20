package elements

// The ExternalDirectoryObjectId (FederatedDirectoryIdentityDetailsType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/externaldirectoryobjectid-federateddirectoryidentitydetailstype
import "encoding/xml"

type ExternalDirectoryObjectIdFederatedDirectoryIdentityDetailsType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
