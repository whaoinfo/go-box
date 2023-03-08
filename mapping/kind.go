package mapping

import "encoding/binary"

const (
	UINT8Size  = 1
	INT8Size   = 1
	UINT16Size = UINT8Size * 2
	INT16Size  = INT8Size * 2
	UINT32Size = UINT16Size * 2
	INT32Size  = INT16Size * 2
	UINT64Size = UINT32Size * 2
	INT64Size  = INT32Size * 2
)

type BytesToNumFuncType func(bytes []byte) int
type NumToBytesFuncType func(num int, bytes []byte)

var (
	bytesToNumFuncMap = map[int]BytesToNumFuncType{
		UINT8Size: func(bytes []byte) int {
			return int(bytes[0])
		},
		UINT16Size: func(bytes []byte) int {
			return int(binary.BigEndian.Uint16(bytes))
		},
		UINT32Size: func(bytes []byte) int {
			return int(binary.BigEndian.Uint32(bytes))
		},
		UINT64Size: func(bytes []byte) int {
			return int(binary.BigEndian.Uint64(bytes))
		},
	}

	numToBytesFuncMap = map[int]NumToBytesFuncType{
		UINT8Size: func(num int, bytes []byte) {
			bytes[0] = byte(num)
		},
		UINT16Size: func(num int, bytes []byte) {
			binary.BigEndian.PutUint16(bytes, uint16(num))
		},
		UINT32Size: func(num int, bytes []byte) {
			binary.BigEndian.PutUint32(bytes, uint32(num))
		},
		UINT64Size: func(num int, bytes []byte) {
			binary.BigEndian.PutUint64(bytes, uint64(num))
		},
	}
)

func GetBytesToNumberFunc(numType int) BytesToNumFuncType {
	return bytesToNumFuncMap[numType]
}

func GetNumberToBytesFunc(numType int) NumToBytesFuncType {
	return numToBytesFuncMap[numType]
}
