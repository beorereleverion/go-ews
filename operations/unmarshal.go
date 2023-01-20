package operations

import (
	"encoding/xml"
	"errors"
	"reflect"
)

func Unmarshal(bArr []byte, e Element) error {
	eAddrVal := reflect.ValueOf(e)
	if eAddrVal.Kind() != reflect.Ptr {
		return &ErrorUnMarshalNotPTR
	}
	eVal := eAddrVal.Elem()
	if !eVal.CanSet() {
		return errors.New("can't set")
	}
	eType := eVal.Type()

	var sfsBA []reflect.StructField
	sfsBA = append(sfsBA, reflect.StructField{
		Name: "Body",
		Type: eType,
		Tag:  reflect.StructTag(`xml:"` + eType.Name() + `"`),
	})
	stBA := reflect.StructOf(sfsBA)
	soBA := reflect.New(stBA).Interface()

	var sfsEA []reflect.StructField
	sfsEA = append(sfsEA, reflect.StructField{
		Name: "Body",
		Type: reflect.TypeOf(soBA),
		Tag:  reflect.StructTag(`xml:"Body"`),
	})
	sfsEA = append(sfsEA, reflect.StructField{
		Name: "XMLName",
		Type: reflect.TypeOf(xml.Name{}),
		Tag:  `xml:"Envelope"`,
	})
	stEA := reflect.StructOf(sfsEA)
	soEA := reflect.New(stEA).Interface()

	err := xml.Unmarshal(bArr, &soEA)
	if err != nil {
		return ErrorCantUnMarshal.formatError(string(bArr), err.Error())
	}

	valArr := reflect.ValueOf(soEA)
	valBo := valArr.Elem().Field(0)
	valElU := valBo.Elem().Field(0)
	eVal.Set(valElU)
	return nil
}
