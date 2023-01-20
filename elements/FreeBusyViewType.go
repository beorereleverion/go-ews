package elements

// The FreeBusyViewType element represents the type of free/busy information returned in the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/freebusyviewtype
import "encoding/xml"

type FreeBusyViewType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Represents the legacy status information: free, busy, tentative, and OOF; the start/end times of the appointments; and various properties of the appointment such as subject, location, and importance. This requested view will return the maximum amount of information for which the requesting user is privileged. If merged free/busy information only is available, as with requesting information for users in a Microsoft Exchange Server 2003 forest, MergedOnly will be returned. Otherwise, FreeBusy or Detailed will be returned.   If Detailed is specified for a distribution list, the free/busy information for the members of the list is merged, and MergedOnly is returned.
	FreeBusyViewTypeDetailed string = `Detailed`
	// Represents all the properties in Detailed with a stream of merged free/busy availability information. If only merged free/busy information is available, for example if the mailbox exists on a computer running Exchange 2003, MergedOnly will be returned. Otherwise, FreeBusyMerged or DetailedMerged will be returned.
	FreeBusyViewTypeDetailedMerged string = `DetailedMerged`
	// Represents the legacy status information: free, busy, tentative, and OOF. This also includes the start/end times of the appointments. This view is richer than the legacy free/busy view because individual meeting start and end times are provided instead of an aggregated free/busy stream.
	FreeBusyViewTypeFreeBusy string = `FreeBusy`
	// Represents all the properties in FreeBusy with a stream of merged free/busy availability information.
	FreeBusyViewTypeFreeBusyMerged string = `FreeBusyMerged`
	// Represents an aggregated free/busy stream. In cross-forest scenarios in which the target user in one forest does not have an Availability service configured, the Availability service of the requestor retrieves the target user's free/busy information from the free/busy public folder. Because public folders only store free/busy information in merged form, MergedOnly is the only available information.
	FreeBusyViewTypeMergedOnly string = `MergedOnly`
	// This value is not valid for requests. This value is valid for responses.
	FreeBusyViewTypeNone string = `None`
)

func (F *FreeBusyViewType) SetForMarshal() {
	F.XMLName.Local = "t:FreeBusyViewType"
}

func (F *FreeBusyViewType) GetSchema() *Schema {
	return &SchemaTypes
}
