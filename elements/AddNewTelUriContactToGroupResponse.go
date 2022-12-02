package elements

// The AddNewTelUriContactToGroupResponse element specifies the result data for the AddNewTelUriContactToGroup WSDL operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addnewteluricontacttogroupresponse
type AddNewTelUriContactToGroupResponse struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The Persona element specifies a set of persona data returned by a GetPersona request.
	Persona *Persona `xml:"m:Persona"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
}
