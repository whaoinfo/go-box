package gtls

import (
	"encoding/binary"
)

type ServerHello struct {
	TLSHello
	CipherSuite       uint16
	CompressionMethod uint8
	extensions        *HelloExtensions
}

func (t *ServerHello) Encode(data []byte) error {
	// random
	offset := RandomLen
	// session id
	t.SessionIDLen = data[offset]
	offset = offset + 1
	sessIdEndIdx := offset + int(t.SessionIDLen)
	t.SessionID = data[offset:sessIdEndIdx]
	offset = offset + int(t.SessionIDLen)

	// cipher suite
	t.CipherSuite = binary.BigEndian.Uint16(data[offset:])
	offset = offset + 2

	// compression method
	t.CompressionMethod = data[offset]
	offset = offset + 1

	// extensions
	t.extensions = &HelloExtensions{}
	return t.extensions.Encode(data[offset:])
}

func (t *ServerHello) FindExtension(extensionType uint16) *HelloExtensionElement {
	if t.extensions == nil {
		return nil
	}

	return t.extensions.Find(extensionType)
}
