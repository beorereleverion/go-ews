package elements

// The PictureData element contains the stream of picture data.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/picturedata
import "encoding/xml"

type PictureData struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PictureData) SetForMarshal() {
	P.XMLName.Local = "m:PictureData"
}

func (P *PictureData) GetSchema() *Schema {
	return &SchemaMessages
}
