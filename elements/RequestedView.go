package elements

// The RequestedView element defines the type of calendar information that a client requests.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/requestedview
type RequestedView string

const (
	// Represents the legacy status information: free, busy, tentative, and OOF; the start/end times of the appointments; and various properties of the appointment such as subject, location, and importance. This requested view will return the maximum amount of information for which the requesting user is privileged. If merged free/busy information only is available, as with requesting information for users in a Microsoft Exchange Server 2003 forest, MergedOnly will be returned. Otherwise, FreeBusy or Detailed will be returned.
	RequestedViewDetailed RequestedView = `Detailed`
	// Represents all the properties in Detailed with a stream of merged free/busy information. If merged free/busy information only is available, MergedOnly will be returned. Otherwise, FreeBusyMerged or DetailedMerged will be returned.
	RequestedViewDetailedMerged RequestedView = `DetailedMerged`
	// Represents the legacy status information: free, busy, tentative, and OOF. This also includes the start/end times of the appointments. This view is richer than the legacy free/busy view because individual meeting start and end times are provided instead of an aggregated free/busy stream.
	RequestedViewFreeBusy RequestedView = `FreeBusy`
	// Represents all the properties in FreeBusy with a stream of merged free/busy information.
	RequestedViewFreeBusyMerged RequestedView = `FreeBusyMerged`
	// Represents an aggregated free/busy stream. In cross-forest scenarios in which the target user in one forest does not have an Availability service configured, the Availability service of the requestor retrieves the target user's free/busy information from the free/busy public folder. Because public folders only store free/busy information in merged form, MergedOnly is the only available information.
	RequestedViewMergedOnly RequestedView = `MergedOnly`
	// This value is not valid for requests. This value is valid for responses.
	RequestedViewNone RequestedView = `None`
)
