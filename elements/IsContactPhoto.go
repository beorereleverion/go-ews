package elements

// The IsContactPhoto element indicates whether the file attachment is a contact picture.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/iscontactphoto
import "encoding/xml"

type IsContactPhoto struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsContactPhotofalse bool = false
	// true
	IsContactPhototrue bool = true
)

func (I *IsContactPhoto) SetForMarshal() {
	I.XMLName.Local = "t:IsContactPhoto"
}

func (I *IsContactPhoto) GetSchema() *Schema {
	return &SchemaTypes
}
