package nbuffer

import "errors"

//import "bytes"

type BufferObject struct {
	//buf *bytes.Buffer
	buf         []byte
	readOffset  int
	writeOffset int
	//realLength int
}

func NewBufferObject(capacity uint) *BufferObject {
	obj := &BufferObject{
		//buf:      make([]byte, 0, capacity),
		buf: make([]byte, capacity),
	}
	return obj
}

func (t *BufferObject) Release() {
	if t.buf == nil {
		return
	}

	t.buf = []byte{}
}

func (t *BufferObject) GetCapacityLen() int {
	return len(t.buf)
}

func (t *BufferObject) Grow(size int) int {
	newBuf := make([]byte, t.GetCapacityLen()+size)
	copy(newBuf, t.buf)
	//t.readOffset = 0
	//t.writeOffset = 0
	return len(t.buf)
}

func (t *BufferObject) Bytes() []byte {
	return t.buf
}

func (t *BufferObject) Read(length int) []byte {
	ret := t.buf[t.readOffset : t.readOffset+length]
	t.readOffset += length
	return ret
}

func (t *BufferObject) GetFreeBytes() []byte {
	return t.buf[t.writeOffset:]
}

func (t *BufferObject) GetFreeBytesLength() int {
	return len(t.GetFreeBytes())
}

func (t *BufferObject) GetNextReadBytes() []byte {
	return t.buf[t.readOffset:]
}

func (t *BufferObject) Rest() {
	t.readOffset = 0
	t.writeOffset = 0
}

func (t *BufferObject) Write(d []byte) error {
	dataLen := len(d)
	if dataLen <= 0 || (t.GetCapacityLen()-t.writeOffset) < dataLen {
		return errors.New("x")
	}

	copy(t.buf[t.writeOffset:], d)
	t.writeOffset += dataLen
	return nil
}

func (t *BufferObject) WriteBytes(bytes ...byte) {
	bytesLen := len(bytes)
	if bytesLen <= 0 {
		return
	}

	growSize := bytesLen - t.GetFreeBytesLength()
	if growSize > 0 {
		t.Grow(growSize * 2)
	}

	copy(t.buf[t.writeOffset:], bytes)
	t.writeOffset += bytesLen
}

func (t *BufferObject) GetRangeBytes(begIdx, len int) []byte {
	return t.buf[begIdx : begIdx+len]
}

func (t *BufferObject) MoveReadOffset(offsetVal int) {
	t.readOffset += offsetVal
}

func (t *BufferObject) UpdateByteValue(index int, value byte) bool {
	if index <= 0 || index >= t.GetCapacityLen() {
		return false
	}

	t.buf[index] = value
	return true
}

func (t *BufferObject) MoveWriteOffset(offsetVal int) {
	t.writeOffset += offsetVal
}

func (t *BufferObject) GetWriteLength() int {
	return t.writeOffset
}

func (t *BufferObject) GetWriteBytes() []byte {
	return t.buf[:t.writeOffset]
}

func (t *BufferObject) GetNextWriteBytes() []byte {
	return t.buf[t.writeOffset:]
}
