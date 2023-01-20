package elements

// The ResponseMessages element contains the response messages for an Exchange Web Services delegate management request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responsemessages-arrayofdelegateuserresponsemessagetype
import "encoding/xml"

type ResponseMessagesArrayOfDelegateUserResponseMessageType struct {
	XMLName xml.Name

	// The DelegateUserResponseMessageType element contains the response message for a single delegate user.
	DelegateUserResponseMessageType *DelegateUserResponseMessageType `xml:"DelegateUserResponseMessageType"`
}

func (R *ResponseMessagesArrayOfDelegateUserResponseMessageType) SetForMarshal() {
	R.XMLName.Local = "m:ResponseMessages"
}

func (R *ResponseMessagesArrayOfDelegateUserResponseMessageType) GetSchema() *Schema {
	return &SchemaMessages
}
