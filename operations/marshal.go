package operations

import (
	"encoding/xml"
	"reflect"
)

func getPTR[T any](s T) *T { return &s }

var (
	startBytes = []byte(`<?xml version="1.0" encoding="UTF-8"?>`)
)

type EnvelopeRequest struct {
	XMLName  xml.Name
	Soap     *Schema `xml:"xmlns:soap,attr"`
	Type     *Schema `xml:"xmlns:t,attr"`
	Messages *Schema `xml:"xmlns:m,attr"`
	Body     *Body
}

func (e *EnvelopeRequest) SetForMarshal() {
	e.XMLName.Local = `soap:Envelope`
}

type Body struct {
	XMLName xml.Name
	Body    interface{}
}

func (b *Body) SetForMarshal() {
	b.XMLName.Local = `soap:Body`
}

type Element interface {
	SetForMarshal()
}

func NewEnvelopeMarshal(body interface{}, schemas ...*Schema) (*EnvelopeRequest, error) {
	reTagXMLElement(body)
	res := &EnvelopeRequest{
		Soap: getPTR(SchemaSOAP),
		Body: &Body{Body: body},
	}
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

func reTagXMLElement(target interface{}) {
	addrValue := reflect.ValueOf(target)
	kind := addrValue.Kind()
	if kind != reflect.Ptr {
		return
	}
	targetValue := addrValue.Elem()
	if !targetValue.IsValid() {
		return
	}
	targetType := targetValue.Type()

	if targetType.Kind() == reflect.Ptr && !targetValue.IsNil() {
		targetValue = targetValue.Elem()
		if !targetValue.IsValid() {
			return
		}
		targetType = targetValue.Type()
	}
	if targetType.Kind() == reflect.Struct {
		for i := 0; i < targetType.NumField(); i++ {
			fValue := targetValue.Field(i)

			if !fValue.IsValid() {
				continue
			}

			if !fValue.CanAddr() {
				continue
			}

			if !fValue.Addr().CanInterface() {
				continue
			}
			if elIface, ok := fValue.Interface().(Element); ok && !fValue.IsNil() {
				elIface.SetForMarshal()
				fValue.Set(reflect.ValueOf(elIface))
			}
			reTagXMLElement(fValue.Addr().Interface())
		}
	}
	if targetType.Kind() == reflect.Slice {
		s := reflect.ValueOf(target)
		if s.Kind() == reflect.Ptr {
			s = s.Elem()
		}

		for i := 0; i < s.Len(); i++ {
			fValue := s.Index(i)

			if elIface, ok := fValue.Interface().(Element); ok && !fValue.IsNil() {
				elIface.SetForMarshal()
				fValue.Set(reflect.ValueOf(elIface))
			}
			reTagXMLElement(fValue.Addr().Interface())
		}
	}
}

func (e *EnvelopeRequest) GetEnvelopeBytes() ([]byte, error) {
	e.SetForMarshal()
	reTagXMLElement(e)
	res, err := xml.Marshal(e)
	if err != nil {
		return nil, ErrorCantMarshal.formatError(*e, err.Error())
	}
	return append(startBytes, res...), nil
}
