package elements

// The ResponseMessages element contains the response messages for an Exchange Web Services delegate management request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responsemessages-arrayofdelegateuserresponsemessagetype
type ResponseMessagesArrayOfDelegateUserResponseMessageType struct {
	// The DelegateUserResponseMessageType element contains the response message for a single delegate user.
	DelegateUserResponseMessageType *DelegateUserResponseMessageType `xml:"m:DelegateUserResponseMessageType"`
}
