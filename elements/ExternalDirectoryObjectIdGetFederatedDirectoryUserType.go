package elements

// The ExternalDirectoryObjectId (GetFederatedDirectoryUserType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/externaldirectoryobjectid-getfederateddirectoryusertype
import "encoding/xml"

type ExternalDirectoryObjectIdGetFederatedDirectoryUserType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
