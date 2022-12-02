package elements

// The FindPeopleResponseMessage element specifies the response message for a FindPeople request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findpeopleresponsemessage
type FindPeopleResponseMessage struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The People element specifies an array of persona data returned as the result of a FindPeople request.
	People *People `xml:"m:People"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
	// The TotalNumberOfPeopleInView element specifies the total number of personas returned in a FindPeople response.
	TotalNumberOfPeopleInView *TotalNumberOfPeopleInView `xml:"m:TotalNumberOfPeopleInView"`
}
