package elements

// The IsWritable element specifies whether the underlying contact or Active Directory recipient can be written to.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/iswritable
import "encoding/xml"

type IsWritable struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsWritablefalse bool = false
	// true
	IsWritabletrue bool = true
)
