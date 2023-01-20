package elements

// The FreeBusyViewOptions element specifies the type of free/busy information returned in the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/freebusyviewoptions
import "encoding/xml"

type FreeBusyViewOptions struct {
	XMLName xml.Name

	// The MergedFreeBusyIntervalInMinutes element represents the time difference between two successive slots in the FreeBusyMerged view.
	MergedFreeBusyIntervalInMinutes *MergedFreeBusyIntervalInMinutes `xml:"MergedFreeBusyIntervalInMinutes"`
	// The RequestedView element defines the type of calendar information that a client requests.
	RequestedView *RequestedView `xml:"RequestedView"`
	// The TimeWindow element identifies the time span queried for the user availability information.
	TimeWindow *TimeWindow `xml:"TimeWindow"`
}

func (F *FreeBusyViewOptions) SetForMarshal() {
	F.XMLName.Local = "t:FreeBusyViewOptions"
}

func (F *FreeBusyViewOptions) GetSchema() *Schema {
	return &SchemaTypes
}
