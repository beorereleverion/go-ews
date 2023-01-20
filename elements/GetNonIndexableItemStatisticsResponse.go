package elements

// The GetNonIndexableItemStatisticsResponse element specifies the response to a GetNonIndexableItemStatistics request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getnonindexableitemstatisticsresponse
import "encoding/xml"

type GetNonIndexableItemStatisticsResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The GetNonIndexableItemStatistics element specifies a request to retrieve nonindexable item statistics.
	GetNonIndexableItemStatistics *GetNonIndexableItemStatistics `xml:"GetNonIndexableItemStatistics"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
}

func (G *GetNonIndexableItemStatisticsResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetNonIndexableItemStatisticsResponse"
}

func (G *GetNonIndexableItemStatisticsResponse) GetSchema() *Schema {
	return &SchemaMessages
}
