package elements

// The GetUserPhoto element contains the request to get a user's photo.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuserphoto
import "encoding/xml"

type GetUserPhoto struct {
	XMLName xml.Name

	// The Email element identifies the email address of the user whose photo is requested in the GetUserPhoto operation.
	Email *EmailString `xml:"Email"`
	// The SizeRequested element contains the requested photo size for a GetUserPhoto operation.
	SizeRequested *SizeRequested `xml:"SizeRequested"`
}

func (G *GetUserPhoto) SetForMarshal() {
	G.XMLName.Local = "m:GetUserPhoto"
}

func (G *GetUserPhoto) GetSchema() *Schema {
	return &SchemaMessages
}
