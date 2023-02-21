package gtls

import (
	"encoding/binary"
)

type ClientHello struct {
	TLSHello
	CipherSuitesLen       uint16
	CipherSuites          []byte
	CompressionMethodsLen uint8
	CompressionMethods    []byte
	extensions            *HelloExtensions
}

func (t *ClientHello) Encode(data []byte) error {
	// random
	offset := RandomLen
	// session id
	t.SessionIDLen = data[offset]
	offset = offset + 1
	sessIdEndIdx := offset + int(t.SessionIDLen)
	t.SessionID = data[offset:sessIdEndIdx]
	offset = offset + int(t.SessionIDLen)

	// cipher suites
	t.CipherSuitesLen = binary.BigEndian.Uint16(data[offset:])
	offset = offset + 2 + int(t.CipherSuitesLen)

	// compression methods
	t.CompressionMethodsLen = data[offset]
	offset = offset + 1 + int(t.CompressionMethodsLen)

	// extensions
	t.extensions = &HelloExtensions{}
	return t.extensions.Encode(data[offset:])
}

func (t *ClientHello) FindExtension(extensionType uint16) *HelloExtensionElement {
	if t.extensions == nil {
		return nil
	}

	return t.extensions.Find(extensionType)
}
