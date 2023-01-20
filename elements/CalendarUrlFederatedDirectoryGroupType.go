package elements

// The CalendarUrl (FederatedDirectoryGroupType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendarurl-federateddirectorygrouptype
import "encoding/xml"

type CalendarUrlFederatedDirectoryGroupType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
