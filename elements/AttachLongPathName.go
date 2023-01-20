package elements

// The AttachLongPathName element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attachlongpathname
import "encoding/xml"

type AttachLongPathName struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
