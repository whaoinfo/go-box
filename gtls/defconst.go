package gtls

type RecordType uint8
type ReadRecordFunc func(data []byte)
type ITLSLayer interface {
	Encode(record []byte) error
}

const (
	RecordTypeChangeCipherSpec RecordType = 20
	RecordTypeAlert            RecordType = 21
	RecordTypeHandshake        RecordType = 22
	RecordTypeApplicationData  RecordType = 23
)

const (
	TypeClientHello uint8 = 1
	TypeServerHello uint8 = 2
)

const (
	RecordLayerHdrSize = 1 + 2 + 2
	HandShakeLenSize   = 3
	HandShakeHdrSize   = 1 + HandShakeLenSize + 2
	RandomLen          = 32
	ExtensionTypeSize  = 2
)
