package elements

// The LegacyDn (NonEmptyArrayOfLegacyDNsType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/legacydn-nonemptyarrayoflegacydnstype
import "encoding/xml"

type LegacyDnNonEmptyArrayOfLegacyDNsType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
