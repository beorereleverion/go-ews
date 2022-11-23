package operations

import (
	"encoding/xml"

	"github.com/beorereleverion/go-ews/elements"
)

// The GetFolder element defines a request to get a folder from a mailbox in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getfolder
type GetFolder struct {
	// Identifies the properties to get for each folder identified in the FolderIds element.
	FolderShape *elements.FolderShape `xml:"m:FolderShape"`
	// Contains an array of folder identifiers that are used to identify folders to get from a mailbox in the Exchange store.
	FolderIds *elements.FolderIds `xml:"m:FolderIds"`
}

type GetFolderEnvelope struct {
	XMLName  xml.Name       `xml:"soap:Envelope"`
	Soap     Schema         `xml:"xmlns:soap,attr"`
	Type     *Schema        `xml:"xmlns:t,attr"`
	Messages *Schema        `xml:"xmlns:m,attr"`
	Body     *GetFolderBody `xml:"soap:Body"`
}

type GetFolderBody struct {
	GetFolder GetFolder `xml:"m:GetFolder"`
}

var (
	ErrorUnsupportedSchema = Error{
		code: 0,
		err:  "unsupported schema name %v",
	}
	ErrorCantMarshal = Error{
		code: 1,
		err:  "cant marshal %#v, err %v",
	}
	ErrorCantUnMarshal = Error{
		code: 2,
		err:  "cant unmarshal %#v, err %v",
	}
)

func NewGetFolderEnvelope(schemas ...*Schema) (*GetFolderEnvelope, error) {
	res := &GetFolderEnvelope{Soap: SchemaSOAP}
	for _, schema := range schemas {
		switch *schema {
		case SchemaTypes:
			res.Type = schema
		case SchemaMessages:
			res.Messages = schema
		default:
			return nil, ErrorUnsupportedSchema.formatError(*schema)
		}
	}
	if len(schemas) == 0 {
		res.Type = getPTR(SchemaTypes)
		res.Messages = getPTR(SchemaMessages)
	}
	return res, nil
}

func (e *GetFolderEnvelope) GetEnvelopeStruct() interface{} {
	return e
}

func (e *GetFolderEnvelope) GetEnvelopeBytes() ([]byte, error) {
	res, err := xml.Marshal(e)
	if err != nil {
		return nil, ErrorCantMarshal.formatError(*e, err.Error())
	}
	return append(startBytes, res...), nil
}
