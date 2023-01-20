package elements

// The GetPasswordExpirationDateResponse element defines the response to a GetPasswordExpirationDate operation operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getpasswordexpirationdateresponse
import "encoding/xml"

type GetPasswordExpirationDateResponse struct {
	XMLName xml.Name

	// The PasswordExpirationDate element provides the password expiration date for a mailbox account.
	PasswordExpirationDate *PasswordExpirationDate `xml:"PasswordExpirationDate"`
	// Describes the status of the response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

const (
	// Describes a request that is fulfilled.
	GetPasswordExpirationDateResponseSuccess = `Success`
	// Describes a request that was not processed. A warning may be returned if an error occurred while an item in the request was processing and subsequent items could not be processed. The following are examples of sources of warnings:  - The Exchange store is offline during the batch.  - Active Directory Domain Services (AD DS) is offline.  - Mailboxes were moved.  - The message database (MDB) is offline.  - A password is expired.  - A quota has been exceeded.
	GetPasswordExpirationDateResponseWarning = `Warning`
	// Describes a request that cannot be fulfilled. The following are examples of sources of errors:  - Invalid attributes or elements.  - Attributes or elements that are out of range.  - An unknown tag.  - An attribute or element that is not valid in the context.  - An unauthorized access attempt by any client.  - A server-side failure in response to a valid client-side call.    Information about the error can be found in the ResponseCode and MessageText elements.
	GetPasswordExpirationDateResponseError = `Error`
)

func (G *GetPasswordExpirationDateResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetPasswordExpirationDateResponse"
}

func (G *GetPasswordExpirationDateResponse) GetSchema() *Schema {
	return &SchemaMessages
}
