package elements

// The ConvertIdResponseMessage element contains the status and result of a ConvertId operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/convertidresponsemessage
import "encoding/xml"

type ConvertIdResponseMessage struct {
	XMLName xml.Name

	// The AlternateId element describes an identifier to convert in a request and the results of a converted identifier in the response.
	AlternateId *AlternateId `xml:"AlternateId"`
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Describes the status of a ConvertId operation response.The following values are valid for this attribute:- Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

const (
	// Describes a request that is fulfilled.
	ConvertIdResponseMessageSuccess = `Success`
	// Describes a request that was not fully processed or for which an unintended result occurred.
	ConvertIdResponseMessageWarning = `Warning`
	// Describes a request that cannot be fulfilled.The following are examples of sources of errors:  - Invalid attributes or elements  - Attributes or elements that are out of range  - An unknown tag  - An attribute or element that is not valid in the context  - An unauthorized access attempt by any client  - A server-side failure in response to a valid client-side callInformation about the error can be found in the ResponseCode and MessageText elements.
	ConvertIdResponseMessageError = `Error`
)

func (C *ConvertIdResponseMessage) SetForMarshal() {
	C.XMLName.Local = "m:ConvertIdResponseMessage"
}

func (C *ConvertIdResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
