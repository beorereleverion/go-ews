package elements

// The GetServerTimeZonesResponseMessage element contains the status and result of a single GetServerTimeZones operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getservertimezonesresponsemessage
type GetServerTimeZonesResponseMessage struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
	// The TimeZoneDefinitions element represents an array of time zone definitions.
	TimeZoneDefinitions *TimeZoneDefinitions `xml:"m:TimeZoneDefinitions"`
	// Describes the status of the response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

const (
	// Describes a request that is fulfilled.
	GetServerTimeZonesResponseMessageSuccess = `Success`
	// Describes a request that was not processed. A warning may be returned if an error occurred while an item in the request was processing and subsequent items could not be processed. The following are examples of sources of warnings:  - The Exchange store is offline during the batch.  - Active Directory Domain Services (AD DS) is offline.  - Mailboxes were moved.  - The message database (MDB) is offline.  - A password is expired.  - A quota has been exceeded.
	GetServerTimeZonesResponseMessageWarning = `Warning`
	// Describes a request that cannot be fulfilled. The following are examples of sources of errors:  - Invalid attributes or elements  - Attributes or elements out of range  - An unknown tag  - An attribute or element that is not valid in the context  - An unauthorized access attempt by any client  - A server-side failure in response to a valid client-side call    Information about the error can be found in the ResponseCode and MessageText elements.
	GetServerTimeZonesResponseMessageError = `Error`
)
