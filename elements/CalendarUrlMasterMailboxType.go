package elements

// The CalendarUrl (MasterMailboxType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendarurl-mastermailboxtype
import "encoding/xml"

type CalendarUrlMasterMailboxType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
