package operations

func getPTR[T any](s T) *T { return &s }

var (
	startBytes = []byte(`<?xml version="1.0" encoding="UTF-8"?>`)
)

// func UnmarshalEnvelope(bb []byte,i interface{})error{
// 	switch v := i.(type) {
// 	case GetFolder:
// 	}

// }
