package elements

// The FreeBusyViewOptions element specifies the type of free/busy information returned in the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/freebusyviewoptions
type FreeBusyViewOptions struct {
	// The MergedFreeBusyIntervalInMinutes element represents the time difference between two successive slots in the FreeBusyMerged view.
	MergedFreeBusyIntervalInMinutes *MergedFreeBusyIntervalInMinutes `xml:"t:MergedFreeBusyIntervalInMinutes"`
	// The RequestedView element defines the type of calendar information that a client requests.
	RequestedView *RequestedView `xml:"t:RequestedView"`
	// The TimeWindow element identifies the time span queried for the user availability information.
	TimeWindow *TimeWindow `xml:"t:TimeWindow"`
}
