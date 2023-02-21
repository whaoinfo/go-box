package gtls

type RecordLayer struct {
	ContentType RecordType
	Version     uint16
	Length      uint16
	Data        []byte

	NextLayer ITLSLayer
}

func (t *RecordLayer) Encode(data []byte) error {
	t.ContentType = RecordType(data[0])
	switch t.ContentType {
	case RecordTypeHandshake:
		t.NextLayer = &TLSHandShakeProtocol{}
		break
	default:
		return nil
	}

	return t.NextLayer.Encode(data[RecordLayerHdrSize:])
}

type TLSHandShakeProtocol struct {
	HandShakeType uint8
	Length        [HandShakeLenSize]byte
	Version       uint16
	Data          []byte

	NextLayer ITLSLayer
}

func (t *TLSHandShakeProtocol) Encode(data []byte) error {
	t.HandShakeType = data[0]
	switch t.HandShakeType {
	case TypeClientHello:
		t.NextLayer = &ClientHello{}
		break
	case TypeServerHello:
		t.NextLayer = &ServerHello{}
		break
	default:
		return nil
	}

	return t.NextLayer.Encode(data[HandShakeHdrSize:])
}

type TLSHello struct {
	Random       [RandomLen]byte
	SessionIDLen uint8
	SessionID    []byte
}
