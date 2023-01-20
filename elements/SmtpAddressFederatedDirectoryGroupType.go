package elements

// The SmtpAddress (FederatedDirectoryGroupType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/smtpaddress-federateddirectorygrouptype
import "encoding/xml"

type SmtpAddressFederatedDirectoryGroupType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
