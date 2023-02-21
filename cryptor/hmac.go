package cryptor

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	hmacKeyElemSize = 2
)

func ComputeHmacSha256(message []byte, secret []byte) (string, string) {
	h := hmac.New(sha256.New, secret)
	h.Write(message)
	sha := hex.EncodeToString(h.Sum(nil))
	return sha, base64.StdEncoding.EncodeToString([]byte(sha))
}

func StringToHmacKey(keyStr string) ([]byte, error) {
	if keyStr == "" {
		return nil, errors.New("key is empty")
	}

	keyElemIdx := 0
	keyElem := make([]string, hmacKeyElemSize)
	keyLen := (len(keyStr) / hmacKeyElemSize) + (len(keyStr) % hmacKeyElemSize)
	retKey := make([]byte, 0, keyLen)

	for _, k := range strings.Split(keyStr, "") {
		keyElem[keyElemIdx] = k
		keyElemIdx = (keyElemIdx + 1) % hmacKeyElemSize
		if keyElemIdx == 0 {
			hexVal, parseErr := strconv.ParseUint(strings.Join(keyElem, ""), 16, 8)
			if parseErr != nil {
				return nil, parseErr
			}

			retKey = append(retKey, byte(hexVal))
		}
	}

	return retKey, nil
}

func HmacKeyToString(hmacKey []byte) string {
	keyLen := len(hmacKey)
	if keyLen <= 0 {
		return ""
	}

	keyElemList := make([]string, 0, keyLen*hmacKeyElemSize)
	fmtStr := "%x"
	for _, elem := range hmacKey {
		if elem < 16 {
			fmtStr = "%x0"
		}
		keyElemList = append(keyElemList, fmt.Sprintf(fmtStr, elem))
	}

	return strings.Join(keyElemList, "")
}
