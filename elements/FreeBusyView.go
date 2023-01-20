package elements

// The FreeBusyView element contains availability information for a specific user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/freebusyview
import "encoding/xml"

type FreeBusyView struct {
	XMLName xml.Name

	// The CalendarEventArray element contains a set of unique calendar item occurrences that represent the requested user's availability.
	CalendarEventArray *CalendarEventArray `xml:"CalendarEventArray"`
	// The FreeBusyViewType element represents the type of free/busy information returned in the response.
	FreeBusyViewType *FreeBusyViewType `xml:"FreeBusyViewType"`
	// The MergedFreeBusy element contains the merged free/busy stream of data.
	MergedFreeBusy *MergedFreeBusy `xml:"MergedFreeBusy"`
	// The WorkingHours element represents the time zone settings and working hours for the requested mailbox user.
	WorkingHours *WorkingHours `xml:"WorkingHours"`
}

func (F *FreeBusyView) SetForMarshal() {
	F.XMLName.Local = "t:FreeBusyView"
}

func (F *FreeBusyView) GetSchema() *Schema {
	return &SchemaTypes
}
