package gtls

import (
	"encoding/binary"
)

type HelloExtensionElement struct {
	ExtType uint16
	Len     uint16
	Data    []byte
}

func (t *HelloExtensionElement) GetExtensionType() uint16 {
	return t.ExtType
}

func (t *HelloExtensionElement) GetExtensionData() []byte {
	return t.Data
}

type HelloExtensions struct {
	Length   uint16
	Elements []*HelloExtensionElement
}

func (t *HelloExtensions) Encode(data []byte) error {
	t.Length = binary.BigEndian.Uint16(data)
	data = data[ExtensionTypeSize:]
	offset := 0
	for offset < int(t.Length) {
		elem := &HelloExtensionElement{}
		elem.ExtType = binary.BigEndian.Uint16(data[offset:])
		offset += ExtensionTypeSize
		elem.Len = binary.BigEndian.Uint16(data[offset:])
		offset += ExtensionTypeSize
		dataEndIdx := offset + int(elem.Len)
		elem.Data = data[offset:dataEndIdx]
		offset += int(elem.Len)
		t.Elements = append(t.Elements, elem)
		// debug
		//log.Printf("Encode Extension, type: %v, length: %v\n", elem.ExtType, elem.Len)
	}

	return nil
}

func (t *HelloExtensions) Find(extensionType uint16) *HelloExtensionElement {
	for _, elem := range t.Elements {
		if elem.ExtType == extensionType {
			return elem
		}
	}

	return nil
}
