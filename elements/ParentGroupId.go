package elements

// The ParentGroupId element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/parentgroupid
import "encoding/xml"

type ParentGroupId struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
