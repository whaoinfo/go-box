package gtls

import (
	"fmt"
	"strings"
)

func SessionID2String(sessionID []byte) string {
	sessIDLen := len(sessionID)
	if sessIDLen <= 0 {
		return ""
	}

	sessionIDBytes := make([]string, 0, sessIDLen*2)
	for _, d := range sessionID {
		fmtStr := "%x"
		if d < 16 {
			fmtStr = "%0x"
		}
		sessionIDBytes = append(sessionIDBytes, fmt.Sprintf(fmtStr, d))
	}
	return strings.Join(sessionIDBytes, "")
}

func NewSessionContext(sessID string) *SessionContext {
	return &SessionContext{
		sessID:       sessID,
		extensionMap: make(map[uint16][]byte),
	}
}

type SessionContext struct {
	sessID       string
	extensionMap map[uint16][]byte
}

func (t *SessionContext) SetExtensionMap(extMap map[uint16][]byte) {
	t.extensionMap = extMap
}

func (t *SessionContext) GetExtensionMap() map[uint16][]byte {
	return t.extensionMap
}
